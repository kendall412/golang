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

func writeToFile(f *os.File, text string) *os.File {
	f.WriteString(text)
	// data, err := io.ReadAll(f)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(data))
	return f
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
	// fmt.Println(f)

	text := "i love march\n"
	x := writeToFile(f, text)

	f1, _ := io.ReadAll(x)
	fmt.Println(string(f1))
}
