
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