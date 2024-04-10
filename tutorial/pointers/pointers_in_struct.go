package main

import "fmt"

/*
pointers in struct
*/
type Hur_ptr struct {
	Fname *string
	Lname *string
	Age   *int
}

type Hur struct {
	Fname string
	Lname string
	Age   int
}

var fname = "naami"
var lname = "hur"
var age = 9

var shinoo = Hur{"shinoo", "hur", 11}
var shinoo_ptr = Hur_ptr{&fname, &lname, &age}

var naami = Hur{"naami", "hur", 9}
var naami_ptr = Hur_ptr{&naami.Fname, &naami.Lname, &naami.Age}

func main() {
	fmt.Println(shinoo)
	fmt.Println()
	// dereference a struct with pointer
	fmt.Println(*shinoo_ptr.Fname)
	fmt.Println(shinoo_ptr.Fname)
	fmt.Println(*shinoo_ptr.Age)
	fmt.Println()
	// dereference a struct with pointer
	fmt.Println(naami_ptr)
	fmt.Println(*naami_ptr.Fname, *naami_ptr.Lname)
}
