package main

import "fmt"

func ret(a,b int) (int,int) {
	aa := a*2
	bb := b*2
	return aa,bb
}

func main() {
	sum1,sum2 := ret(1,3)
	//fmt.Println("sum1 ",sum1)
	fmt.Println("sum1 of 1,3 is ",sum1,sum2)
}
