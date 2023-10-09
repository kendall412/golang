package main

import (
	"fmt"
)

func main(){
	// in go we don't have exceptions like other lang
	// a lot of exceptions in other languages are considered
	// normal in golang
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}