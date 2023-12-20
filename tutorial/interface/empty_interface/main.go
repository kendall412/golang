package main

import "math"

type nointer interface{}

type rect struct{
	width float64
	height float64
}

type circle struct{
	radius float64
}

func (r rect) area()float64{
	return 2*r.width*r.height
}

func (c circle) area() float64{
	return math.Pi*c.radius*c.radius
}

func (nointer interface) getArea()