package main

import "fmt"

type ninja struct {
	name string
	weapons []string
	level int
}

func main() {
	wallace := ninja{name: "Wallace"}
	wallace = ninja{
		name: "Wallace",
		weapons: []string{"Ninja Start","Ninja Sword"},
		level: 1,
	}
	fmt.Println(wallace.name)
	fmt.Println(wallace.weapons)
	fmt.Println(wallace.level)
	wallace.level++
	fmt.Println(wallace.level)
	wallace.weapons = append(wallace.weapons, "Sai")
	fmt.Println(wallace)
}