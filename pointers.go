package main

import (
	"fmt"
)

// func main() {
// 	var a int = 42
// 	var b *int = &a
// 	// pointer to an interger
// 	fmt.Println(a, *b)
// 	// *b is dereferencing meaning getting the value at the
// 	// address it is pointing to
// 	*b = 14
// 	fmt.Println(a, *b)
// }

func main() {
	
	a := [...]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	// var x *int = &a[0]
	// var y *int = &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
	// fmt.Printf("%p %p\n", *x, *y)
	// %p prints the value of the pointer
}