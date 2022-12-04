
// program to find the area of a rectangle

package main

import ("fmt"
	    "math")

type Rectangle struct{
	length float64
	width float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) area() float64{
	return r.length*r.width
}

func (c Circle) area() float64{
	return math.Pi*c.radius*c.radius