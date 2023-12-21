/*
GENERAL:

Width describes the number of bytes of storage an instance of a type occupies. As a process's address space is one dimensional, width is more apt than size.

Width is property of a type. As evey value in a Go program has a type, the width of the value is defined by its type and is always a multiple of 8 bits.

Width can be discovered by using unsafe.Sizeof() function.

e.g.:
var c string
fmt.Println(unsafe.Sizeof(c)) // prints 8

var [a]unint32
fmt.Println(unsafe.Sizeof(a)) // prints 12

type S struct{
	a uint16
	b uint32
}
var s S
fmt.Println(unsafe.Sizeof(s)) // prints 8 not 6

EMPTY STRUCT:

Empty structs occupy zero width and zero bytes of storage. Thus it follows that it does not need padding.

e.g.:
var s [1000000]struct{}
fmt.Println(unsafe.Sizeof(s)) // prints 0

slices of struct{} consume only the space for their slice headers.

var x = make([]struct{},1000000)
fmt.Println(unsafe.Sizeof(x)) // prints 12

Empty structs behave just like any other structs. Therefore, empty structs can be used as method receivers.

e.g.:

type S struct{}

func (s *S) add() {
	fmt.Printf("%p\n",s)
}

func main(){
	var a,b S
	a.add() // gives memory address
	b.add() // gives memory address
}

https://dave.cheney.net/2014/03/25/the-empty-struct
*/

package main

import "fmt"

type S struct{}

func (s *S) addr() {
	fmt.Printf("%p\n", s)
}

func main() {
	var a S
	a.addr()
}
