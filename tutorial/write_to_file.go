package main

import (
	"fmt"
	"io"
	"os"
)

type newfiles struct {
	name     string
	ext      string
	filetype string
}

func createFile(f newfiles) {
	file, err := os.Create(f.name + f.ext)
	if err != nil {
		fmt.Printf("Could not create file %s\n", f)
		fmt.Println(err)
	}
	text := "I love March.\nI love my kids!"
	file.WriteString(text)
	defer file.Close()
}

// creates a file (txt, log) and returns *os.File
func createAndreturnFile(f newfiles) *os.File {
	filename := f.name + f.ext
	/*
		os.Create()
			Create creates or truncates the named file. If the file already exists, it is truncated. If the file does not exist, it is created with mode 0666 (before umask). If successful, methods on the returned File can be used for I/O; the associated file descriptor has mode O_RDWR. If there is an error, it will be of type *PathError.
	*/
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Could not create file %s\n", f)
		fmt.Println(err)
	}
	return file
}

func readFile(file newfiles) {
	f, err := os.ReadFile(file.name + file.ext)
	if err != nil {
		fmt.Printf("could not read file %s\n", file)
	}
	f_str := string(f)
	fmt.Println(f_str)
}

func writeToFile(f *os.File, text string) {
	f.WriteString(text)
	// return f
}

func openFile() {
	f, err := os.Open("newfile.txt")
	// The open method returns a pointer to the named file, the returned file can
	// be used for reading; and the associated file descriptor. The retured file
	// can now be used in I/O operations.
	// https://medium.com/@kdnotes/exporing-files-and-folder-from-the-golang-os-pacckage-a8984a8bf055

	if err != nil {
		fmt.Println("could not read file")
	}
	data, _ := io.ReadAll(f)
	fmt.Println(string(data))
}

func main() {
	n1 := newfiles{name: "newfile", ext: ".txt", filetype: "text"}
	var f *os.File
	f = createAndreturnFile(n1)
	fmt.Printf("f type: %T\n", f)

	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.\nUt enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat\nDuis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.\nSed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo\nNemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur\nQuis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?\n"
	writeToFile(f, text)

	nf, _ := os.Open(n1.name + n1.ext)
	f1, _ := io.ReadAll(nf)
	fmt.Println(string(f1))
}
