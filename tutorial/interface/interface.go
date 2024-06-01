package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

type shape interface {
	area() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	// var s1 shape
	// s1 = circle{10}
	// fmt.Println(s1.area())

	// var s2 shape
	// s2 = rect{2, 3}
	// fmt.Println(s2.area())

	// var s3 shape
	// s3 = circle{95.3}
	// fmt.Println(getArea(s3))

	c1 := circle{4.5}
	r1 := rect{5, 7}
	shapes := []shape{c1, r1}
	fmt.Println(shapes[0])
	fmt.Println(shapes[1])
	fmt.Println(shapes[0].area())
	fmt.Println(shapes[1].area())

	for index, shape := range shapes {
		fmt.Printf("index: %b area: %f\n", index, getArea(shape))
	}

}
