
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