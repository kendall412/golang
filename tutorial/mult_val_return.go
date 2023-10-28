package main

import (
	f "fmt"
)

func ret(a, b int) (int, int) {
	aa := a * 2
	bb := b * 2
	return aa, bb
}

func main() {
	val1 := 1
	val2 := 3
	result1, result2 := ret(val1, val2)
	f.Printf("ret(%d, %d): %d, %d\n", val1, val2, result1, result2)
}
