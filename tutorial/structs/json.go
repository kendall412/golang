package main

import (
	"encoding/json"
	"fmt"
	// "encoding/json"
)

type Book struct {
	Title  string
	Author string
}

type T struct {
	F1 string `json:"field1"`
	F2 string `json:"field2,omitempty"`
	F3 string `json:"field3,omitempty"`
	F4 string `json:"-"`
}

func main() {
	// type Book struct
	book := Book{
		Title:  "Learning Concurrency in Python",
		Author: "Elliot Forbe",
	}
	fmt.Printf("%+v\n", book)

	// type T struct
	t := T{
		F1: "v1",
		F2: " ",
		F3: "v3",
		F4: "v4",
	}

	s, _ := json.Marshal(t)
	fmt.Printf("%s\n", s)
}
