
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reading data from a file
	file, err := os.Open("equation.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	// reading the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		//parse the text
		s := strings.Split(text, " ")
		a, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Println(err)
		}
		b, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println(err)
		}
		c, err := strconv.Atoi(s[2])
		if err != nil {
			fmt.Println(err)
		}
		//calculating the equation
		fmt.Println(solve(a, b, c))
	}
}
func solve(a, b, c int) ([2]float64, error) {
	Delta := math.Pow(float64(b), 2) - 4*float64(a)*float64(c)
	if Delta < 0 {
		return [2]float64{}, fmt.Errorf("Delta is less than zero, the equation has no solution")