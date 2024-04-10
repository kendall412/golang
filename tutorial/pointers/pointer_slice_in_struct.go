package main

import "fmt"

type Persons struct {
	persons       []*Hur
	family_status string
}

type Hur struct {
	Fname string
	Lname string
	Age   int
}

var shinoo = Hur{"shinoo", "hur", 11}
var naami = Hur{"naami", "hur", 9}
var woojin = Hur{"woojin", "hur", 6}
var wonoo = Hur{"wonoo", "hur", 4}

var pers1 = Person{
	family_status: "nil",
	persons:       Hur{&shinoo, &naami, &woojin, &wonoo},
}

func main() {
	fmt.Println(pers1)

}
