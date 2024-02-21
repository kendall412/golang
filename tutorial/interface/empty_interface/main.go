/*
Empty Interface:

Normally in functions, when we pass values to the function parameters, we need to specify the data type of parameters in a function definition. However, with an empty interface, we can pass parameters of any data type.
We can also use an empty interface to pass any number of arguments to the function definition.

https://www.programiz.com/golang/empty-interface
*/

package main

import "fmt"

// 1. empty interface declared prior to the function
type empty interface{}

func say(a ...empty) {
	for _, v := range a {
		fmt.Println(v)
	}
}

// 2. function that receives variadic empty interface
func say2(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	x := "danny"
	y := 51
	z := true
	t := []string{"shinoo", "naami", "woojin", "wonoo"}
	s := map[string]string{
		"one":   "one1",
		"two":   "two2",
		"three": "three3",
	}
	say(x, y, z, s)
	fmt.Println()

	say2(x, t)
}
