package main

import (
	"fmt"
)

func main() {
	i, j := 2, 10
	fmt.Println(i, j)
	fmt.Println(&i, &j)
	// for every '&' say address of
	// &i is address of i
	p := &i
	// assign p to the address of i
	fmt.Println(p)
	fmt.Println(*p)
	// * has 2 uses:
	// if * is in front of a type (e.g. *int) this whole thing is a type, or a pointer type with int as a base
	// if * is in front of a variable (e.g. *p) the * acts
	// as an operator (dereferencing)
	
}
