package main

import (
	"fmt"
)

func main() {
	// arrays starts with a size, type (only 1 type is allowed), and can use
	// initializer by using {...}
	// arrays have fixed size
	// grades := [...]int{97, 85, 93}
	var students [3]string
	// fmt.Printf("Grades: %v", grades)
	fmt.Printf("Students: %s\n",students )
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %s\n",students)
}