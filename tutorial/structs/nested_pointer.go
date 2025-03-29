package main

import (
	// abbreviating import module names
	f "fmt"
)

type Person struct {
	Fname string
	Age   int
}

type Hur struct {
	person   Person
	Lname    string
	Position string
}

// pointer struct
func (p *Person) changePerson(fname string, age int) {
	p.Fname = fname
	p.Age = age
}

func (h *Hur) changeHur(lname, position string) {
	h.Lname = lname
	h.Position = position
}

func main() {
	// struct pointer
	person := &Person{
		Fname: "Shinoo",
		Age:   10,
	}

	hur := &Hur{
		person: Person{
			Fname: "Naami",
			Age:   8,
		},
		Lname:    "Hur",
		Position: "daughter",
	}

	f.Println(person.Fname)
	f.Println(person.Age)
	f.Println()
	f.Println(hur.person.Fname)
	f.Println(hur.Lname)
	f.Println(hur.person.Age)
	f.Println(hur.Position)
	f.Println()
	hur.changeHur("Hom", "single")
	f.Println(hur.person.Fname)
	f.Println(hur.Lname)
	f.Println(hur.person.Age)
	f.Println(hur.Position)
}
