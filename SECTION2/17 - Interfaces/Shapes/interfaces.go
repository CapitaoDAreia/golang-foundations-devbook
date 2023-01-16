package main

import (
	"fmt"
	"math"
)

type retangle struct {
	width  float64
	height float64
}

func (r retangle) area() float64 {
	/*
		Implements area() in retangle struct.
		This turns possible that writeArea receives retangle
	*/
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	/*
		Implements area() in circle struct.
		This turns possible that writeArea receives circle
	*/
	return math.Pi * math.Pow(c.radius, 2)
}

type form interface {
	area() float64 //interfaces only must have signature of methods
}

func writeArea(f form) {
	fmt.Printf("The area is %0.2f\n", f.area())
}

func main() {
	r := retangle{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
