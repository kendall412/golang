/*
Modify slice by passing slice pointer.
*/
package main

import "fmt"

type Cart struct {
	Id   string
	Paid bool
}

func changeByPointer(cartptr *Cart) {
	*cartptr = Cart{Id: "9999", Paid: false}

	fmt.Println("change() cartptr: ", cartptr)
	fmt.Println("change() *cartptr: ", *cartptr)
	fmt.Println()
}

func main() {
	cart := Cart{Id: "2322", Paid: true}
	fmt.Println("cart: ", cart)
	fmt.Println()
	changeByPointer(&cart)
	fmt.Println("cart (after change(&cart))", cart)
}
