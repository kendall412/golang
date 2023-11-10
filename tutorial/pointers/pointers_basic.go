package main

import (
	f "fmt"
)

func changeName(name *string, nname string) {
	*name = nname
}

func main() {
	name := "Danny"
	f.Println("Original Name: ", name)

	changeName(&name, "March")
	f.Println("Name after applying pointer func: ", name)
}
