package main

import (
	"fmt"
)

func main() {
	// only square brackets, type, and pass initializer
	// slices are reference types (can change original slice by changing it somewhere else)
	a := []int{1,2,3}
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}