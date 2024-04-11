package main

import "fmt"

type Hur struct {
	Fname string
	Lname string
}

type Person struct {
	People Hur
	Age    int
}

type Persons struct {
	Peoples []Hur
	Desc    string
}

func main() {
	var naami = Hur{"naami", "hur"}
	var woojin = Hur{"woojin", "hur"}
	//var shinoo = Hur{"shinoo","hur"}
	var per1 = Person{
		Age: 11,
		People: Hur{
			Fname: "shinoo",
			Lname: "hur"},
	}
	var per2 = Person{
		Age:    9,
		People: naami}

	fmt.Println(per1)
	fmt.Println(per2)

	/*
		slice in struct
	*/
	hurs := []Hur{}
	hurs = append(hurs, woojin, naami)
	fmt.Println(hurs)
	var pers1 = Persons{
		Desc:    "The Hur family",
		Peoples: hurs,
	}
	fmt.Println(pers1)
}
