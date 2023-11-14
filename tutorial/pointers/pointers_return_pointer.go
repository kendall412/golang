package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() *person {
	m := person{name: "noname", age: 50}
	return &m
}

func changeName(name *string) {
	*name = "March"
}

func changeName2(name string) *string {
	name = "March"
	fmt.Println("memory location of name: ", &name)
	return &name
}

func main() {
	// passing pointer
	x := "Danny"
	fmt.Println(x)
	changeName(&x)
	fmt.Println(x)

	// returning pointer
	y := "Danny"
	fmt.Println("memory location of y: ", &y)
	changeName2(y)
	fmt.Println(y)
	fmt.Println("memory location of y: ", &y)
}
