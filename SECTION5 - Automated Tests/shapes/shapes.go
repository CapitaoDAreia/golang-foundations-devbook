package main

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	/*
		Implements area() in Retangle struct.
		This turns possible that writeArea receives Retangle
	*/
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	/*
		Implements area() in circle struct.
		This turns possible that writeArea receives circle
	*/
	return math.Pi * math.Pow(c.Radius, 2)
}

type Form interface {
	Area() float64 //interfaces only must have signature of methods
}
