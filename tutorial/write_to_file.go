package main

import (
	"fmt"
	"os"
)

func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("Could not create file %s\n", name)
		fmt.Println(err)
	}

	text := "i love march\n"
	file.WriteString(text)

	defer file.Close()
}

func readFile(file string) {
	// text := "i love march\n"
	f, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("could not read file %s\n", file)
	}
	fmt.Println(f)
}

func main() {
	fname := "newfile.txt"
	createFile(fname)
	readFile(fname)

}
