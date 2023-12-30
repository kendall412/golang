package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Hur struct {
	Person
	family_position string
}

type Emotional interface {
	cry()
	workhard()
}

func (h Hur) cry() {
	fmt.Println(h.name, " is a cry baby.")
}

func (h Hur) tantrum() {
	fmt.Println(h.name, " is having a tantrum.")
}

func (p Person) cry() {
	fmt.Println(p.name, " also cries too.")
}

func (h Hur) workhard() {
	fmt.Println("Hur family works hard")
}

func (p Person) workhard() {
	fmt.Println(p.name, " works hard.")
}
func main() {
	// var hur_family []Hur

	mike := Person{
		"Mike",
		52,
	}

	// noo := Hur{
	// 	Person: Person{
	// 		"Shinoo",
	// 		10,
	// 	},
	// 	family_position: "eldest",
	// }

	// won := Hur{
	// 	Person: Person{
	// 		"Wonoo",
	// 		4,
	// 	},
	// 	family_position: "youngest",
	// }

	// march := Hur{
	// 	Person: Person{
	// 		"March",
	// 		42,
	// 	},
	// 	family_position: "mom",
	// }

	danny := Hur{
		Person: Person{
			"Danny",
			51,
		},
		family_position: "father",
	}

	// hur_family = append(hur_family, noo, won, march, danny)
	// fmt.Println(hur_family)

	Emotional.cry(danny)
	Emotional.cry(mike)

	var emo Emotional
	emo = mike
	emo.workhard()
}
