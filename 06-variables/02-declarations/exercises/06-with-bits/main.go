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
)

// ---------------------------------------------------------
// EXERCISE: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
//
// 3. After you've done, check out the solution
//    and read the comments there
//
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// ---------------------------------------------------------

func main() {
	var i1 int
	var i2 int8
	var i3 int16
	var i4 int32
	var i5 int64
	var i6 float32
	var i7 float64
	var i8 complex64
	var i9 complex128
	var i10 bool
	var i11 string
	var i12 rune
	var i13 byte

	fmt.Println(
		i1, i2, i3,
		i4, i5, i6,
		i7, i8, i9,
		i10, i11, i12,
		i13)
}

