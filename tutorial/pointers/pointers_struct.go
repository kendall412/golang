package main

import "fmt"

type Engineer struct {
	Name string
	Age  int
}

func changeValue(str *string) {
	*str = "world!"
}

/*
use below when Engineer is instantiated as pointer.
e.g. eng := &Engineer{...} rather than eng := Engineer{...}
*/
func (e *Engineer) changeName() {
	e.Name = "Shinoo"
}

func (e *Engineer) changeAge() {
	e.Age = 42
}

func (e *Engineer) changeValues(new_name string, new_age int) {
	e.Name = new_name
	e.Age = new_age
}

func (e Engineer) changeValues2() (*string, *int) {
	e.Name = "Wonoo"
	e.Age = 4
	return &e.Name, &e.Age
}

func main() {
	// basic pointer
	toChange := "hello"
	fmt.Println("Original toChange: ", toChange)
	changeValue(&toChange)
	fmt.Println("after applying pointer func: ", toChange)

	// pointer in struct
	eng := &Engineer{
		Name: "March",
		Age:  42,
	}

	eng.changeValues("Woojin", 6)
	fmt.Println(eng.Name, eng.Age)
}
