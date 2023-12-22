package main

import "fmt"

type ninja struct {
	name    string
	weapons []string
	level   int
}

type person struct {
	fname string
	lname string
	age   int
}

type persons struct {
	indiv person
	hobby string
}

func main() {
	// using ninja
	wallace := ninja{name: "Wallace"}
	wallace = ninja{
		name:    "Wallace",
		weapons: []string{"Ninja Start", "Ninja Sword"},
		level:   1,
	}
	fmt.Println(wallace.name)
	fmt.Println(wallace.weapons)
	fmt.Println(wallace.level)
	wallace.level++
	fmt.Println(wallace.level)
	wallace.weapons = append(wallace.weapons, "Sai")
	fmt.Println(wallace)

	fmt.Printf("\n\n")

	// Nested struct:: using person and persons
	person1 := persons{
		hobby: "motorcycle",
		indiv: person{"Mike", "Johnson", 41},
	}
	fmt.Printf("person's first name: %s\n", person1.indiv.fname)
	fmt.Printf("person's last name: %s\n", person1.indiv.lname)
	fmt.Printf("person's age: %d\n", person1.indiv.age)
	fmt.Printf("person's hobby: %s\n", person1.hobby)
}
