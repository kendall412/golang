/*
Functions declared inside of functions are special; they are closures. This means that functions declared inside of functions are able to access and modify variables declared in the outer function.
*/

package main

import "fmt"

func main() {
	a := 100
	f := func() {
		fmt.Println(a)
		/*
			Using := instead of = inside the closure creates a new a that ceases to exist when the closure exits.
		*/
		a = 50
		fmt.Println(a)
	}

	f()
	fmt.Println(a)
}
