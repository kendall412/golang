package main

import "fmt"

/*
using pointer of struct in slice
*/
type Persons struct {
	persons       []*Hur
	family_status string
}

type Hur struct {
	Fname string
	Lname string
	Age   int
}

func regularSlice() {
	var shinoo = Hur{"shinoo", "hur", 11}
	var naami = Hur{"naami", "hur", 9}
	var woojin = Hur{"woojin", "hur", 6}
	var wonoo = Hur{"wonoo", "hur", 4}

	hurs1 := []Hur{}
	hurs1 = append(hurs1, shinoo, naami, wonoo, woojin)

	// fmt.Println(pers1)
	fmt.Println(hurs1)
	fmt.Println(hurs1[0])
}

/*
pointer slice in struct
*/
func pointerSlice() {
	shinoo := Hur{"shinoo", "hur", 11}
	naami := Hur{"naami", "hur", 9}
	woojin := Hur{"woojin", "hur", 6}
	wonoo := Hur{"wonoo", "hur", 4}

	var pers1 = Persons{
		family_status: "children",
		persons:       []*Hur{&shinoo, &naami, &woojin, &wonoo},
	}

	fmt.Println(pers1)
	// will print slice of memory addressss of shinoo, naami, woojin, wonoo Hur type
	// {[0xc0000700c0 0xc0000700f0 0xc000070120 0xc000070150] children}

	fmt.Println(pers1.persons)
	// will print memory addresses of persons type (*Hur) of shinoo, naami , woojin, wonoo
	// [0xc0000700c0 0xc0000700f0 0xc000070120 0xc000070150]

	fmt.Println(pers1.family_status)
	// will print family status (string)
	// children

	fmt.Println(*pers1.persons[2])
	// will print deferenced slice element in position 2 {woojin hur 6}
	// {woojin hur 6}

	fmt.Println(pers1.persons[3].Fname)
	// will print first name of slice element in position 3
	// wonoo
}

func main() {
	// regularSlice()
	pointerSlice()
}
