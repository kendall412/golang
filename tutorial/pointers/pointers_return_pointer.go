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

func main() {
	x := "Danny"
	fmt.Println(x)
	changeName(&x)
	fmt.Println(x)
}
