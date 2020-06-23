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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"

	var nums []int
	for _, v := range strings.Fields(data) {
		v, _ := strconv.Atoi(v)
		nums = append(nums, v)
	}

	var evens, odds []int
	for _, v := range nums {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}

	var middle []int
	if l := len(nums); l%2 == 0 {
		middle = nums[l/2-1 : l/2+1]
	} else {
		middle = nums[l/2 : l/2+1]
	}

	first2 := nums[:2]
	last2 := nums[len(nums)-2:]
	evensLast1 := evens[len(evens)-1:]
	oddsLast2 := odds[len(odds)-2:]

	fmt.Printf("nums           : %d\n", nums)
	fmt.Printf("evens          : %d\n", evens)
	fmt.Printf("odds           : %d\n", odds)
	fmt.Printf("middle         : %d\n", middle)
	fmt.Printf("fitst 2        : %d\n", first2)
	fmt.Printf("last 2         : %d\n", last2)
	fmt.Printf("evens last 1   : %d\n", evensLast1)
	fmt.Printf("odds last 2    : %d\n", oddsLast2)

}
