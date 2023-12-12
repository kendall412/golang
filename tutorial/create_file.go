package main

import (
	"fmt"
	"os"
)

func createFile(name, format string) {
	var ext string
	if format == "text" {
		ext = "txt"
	}

	if format == "log" {
		ext = "log"
	}

	filename := name + "." + ext
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Could not create file %s", filename)
	}

	defer file.Close()
}

func main() {
	createFile("newfile", "text")
}
