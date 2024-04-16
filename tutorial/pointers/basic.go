package main

import (
	f "fmt"
	"strings"
)

var border = strings.Repeat("#", 20)

func changeName(name *string, new_name string) {
	f.Println(border)
	*name = new_name
	// f.Printf("name: %p\n", name)
	// f.Printf("name type: %T\n", name)
	// f.Printf("*name: %s\n", *name)
	f.Println(border)
}

func main() {
	name := "Danny"
	f.Println("Original Name: ", name)

	var name_ptr *string
	name_ptr = &name
	new_name := "March"
	// f.Println("new_name: " + new_name)
	changeName(name_ptr, new_name)

	f.Printf("type of name_ptr: %T\n", name_ptr)
	f.Println("Name after applying pointer func: ", name)
}
