package main

import (
	"fmt"
	// "encoding/json"
)

type Book struct {
	Title string
	Author string
}

func main() {
	book := Book{Title: "Learning Concurrency in Python", Author: "Elliot Forbe"}
	fmt.Printf("%+v\n", book)
}