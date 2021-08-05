package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	// moves it after the main function but
	// before main returns
	fmt.Println("end")


	a := "alpha"
	defer fmt.Println(a)
	// will use "alpha"
	// when defer it uses the value of variable at the time
	// the defer was called
	a = "omega"
}

