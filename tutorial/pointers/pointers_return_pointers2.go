package main

import "fmt"

func change(x string) *string {
	s := "Hello " + x
	return &s
}

func main() {
	y := "Danny"
	var msg *string
	msg = change(y)
	fmt.Println(msg, *msg)
}
