package main

import "fmt"

var phrase string

func init(){
	phrase = "Hola Mundo!!"
	fmt.Println("This is the constructor")
	}

func main(){
	fmt.Println (phrase)
	}