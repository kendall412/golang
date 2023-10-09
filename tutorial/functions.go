package main

import (
	"fmt"
)

// uppercase name is public
// lowercase name is internal-private only

func main() {
	for i := 0; i < 5; i++ {	
		sayMessage("Hello Go!", i)
	}
	sayGreeting("Hello", "Stacey")
}

func sayMessage(msg string, idx int){
	fmt.Println(msg)
	fmt.Println("The value of the index is ", idx)
}

// func sayGreeting(greeting string, name string) {
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}