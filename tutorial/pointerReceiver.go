package main

import "fmt"

func main(){
	x := "hello danny"
	y := &x
	fmt.Println(x)
	change(y)
	fmt.Println(x)
}

func change(msg *string){
	*msg = "I like learning golang"
}