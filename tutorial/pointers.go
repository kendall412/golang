package main

import "fmt"

func changeValue(str *string) {
	// *string is pointer type
	// *str is dereference
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	toChange := "hello"

	fmt.Println(toChange)
	changeValue(&toChange) // &toChange is the pointer
	fmt.Println(toChange)

	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
}
