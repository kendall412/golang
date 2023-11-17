package main

import (
	f "fmt"
	"strconv"
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

func (h Hur) create() *string {
	// converts int to string
	age := strconv.Itoa(h.person.Age)
	msg := "Hello " + h.person.Fname + ". You are " + age + " and your last name is " + h.Lname + " and you are " + h.Position + " to Danny."
	return &msg
}

func (h *Hur) changeVal(nfname string, nlname string, nposition string, nage int) {
	h.Lname = nlname
	h.Position = nposition
	h.person.Fname = nfname
	h.person.Age = nage
}

func main() {
	hur := &Hur{
		person: Person{
			Fname: "Shinoo",
			Age:   10,
		},
		Lname:    "Hur",
		Position: "son",
	}

	// // use the return value directly from method invocation
	// f.Println("Directly calling method *hur.create()")
	// f.Println(*hur.create())

	// // use short hand variable declaration
	// f.Println("Short hand var declaration msg := hur.create()")
	// msg := hur.create()
	// f.Println(msg)

	// use long hand variable declaration
	f.Println("::: Long hand var declaration var msg1 *string :::")
	var msg1 *string
	msg1 = hur.create()
	f.Println(*msg1)

	hur.changeVal("Naami", "Hur", "daughter", 8)
	msg1 = hur.create()
	f.Println(*msg1)
}
