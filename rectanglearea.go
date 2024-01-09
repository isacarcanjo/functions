package main

import (
	"fmt"
)

func main() {

	// rect1 := Rectangle{10.0, 10.0}
	// rect2 := Rectangle{20.0, 10.0}
	// rect3 := Rectangle{10.0, 20.0}

	// fmt.Println("Area of rect1 is: ", getArea(rect1))
	// fmt.Println("Area of rect2 is: ", getArea(rect2))
	// fmt.Println("Area of rect3 is: ", getArea(rect3))

	rect := Rectangle{10.0, 10.0}
	fmt.Println("Area of rect is: ", rect.area())

}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

// func getArea(rectangle Rectangle) float64 {
// 	return rectangle.length * rectangle.width
// }




"""@author: github.com/isacarcanjo"""



