package main

import "fmt"

type Hur struct {
	Fname string
	Lname string
}

type Persons struct {
	Peoples []Hur
	Desc    string
}

/*
slice in struct
*/
func main() {
	pers := Persons{
		Desc: "Hur family",
		Peoples: []Hur{
			{Fname: "naami", Lname: "hur"},
			{Fname: "shinoo", Lname: "hur"},
			{Fname: "woojin", Lname: "hur"},
			{"wonoo", "hur"},
		},
	}
	fmt.Println(pers)
	fmt.Println(pers.Desc)
	fmt.Println(pers.Peoples)
	fmt.Println(pers.Peoples[0].Fname)
	fmt.Println(pers.Peoples[3].Fname)
}
