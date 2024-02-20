/*
How to modify slice with pointer.
*/
package main

import "fmt"

type Cart struct {
	Id   string
	Paid bool
}

var cart = Cart{Id: "2322", Paid: true}

func change(cartptr *Cart) {
	*cartptr = Cart{Id: "9999", Paid: false}
}

func main() {
	fmt.Println(cart)
	change(&cart)
	fmt.Println(cart)

}
