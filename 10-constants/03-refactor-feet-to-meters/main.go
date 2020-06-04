// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// EXERCISE
//
// This program uses magic numbers
// Refactor it and use named constants instead

func main() {
	const (
		toMeters = 0.3048
		toYards = 0.3333
	)
	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * toMeters
	yards := math.Round(feet * toYards)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
