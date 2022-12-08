
package main

import (
	"fmt"
	"math"
)

func area(radius float64) (float64) {
	return math.Pi * radius * radius
}

func main() {
	var radius float64
	fmt.Print("Enter radius: ")