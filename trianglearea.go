
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
}

func main(){
	r1 := Rectangle{
		length: 40,
		width: 20,
	}

	c1 := Circle{
		radius: 20,
	}

	fmt.Println("Area of Rectangle is:", r1.area())
	fmt.Println("Area of Circle is:", c1.area())
}


