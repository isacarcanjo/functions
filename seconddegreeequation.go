package main

import (
	"fmt"
	"math"
)

type eq struct {
	a float64
	b float64
	c float64
}

func (e *eq) delta() float64 {
	return e.b*e.b - 4*e.a*e.c
}
func (e *eq) compute() {
	d := e.delta()
	if d < 0 {
		fmt.Println("no solutions")
	} else if d == 0 {
		fmt.Println(-e.b / 2)
	} else {
		fmt.Println((-e.b - math.Sqrt(d)) / 2*e.a)
		fmt.Println((-e.b + math.Sqrt(d)) / 2*e.a)
	}
}

func main() {
	e := eq{1, 2, 1}
	fmt.Println(e.delta())
	e.compute()
}




"""@author: github.com/isacarcanjo"""



