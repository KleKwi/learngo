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
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	_ = strings

	// Print the byte and rune length of the strings
	// Hint: Use len and utf8.RuneCountInString
	for _, s := range strings {
		fmt.Printf("%q\n", s)
		fmt.Printf("\tbyte length: %d\n", len(s))
		fmt.Printf("\trune length: %d\n", utf8.RuneCountInString(s))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes: % [1]x\n", s)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Print("\trunes: ")
		for _, r := range s {
			fmt.Printf("% x", r)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Print("\trunes: ")
		for _, r := range s {
			fmt.Printf("%q ", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst rune: %q, size: %d\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		l, size := utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tlast rune: %q, size: %d\n", l, size)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:])
		fmt.Printf("\tfirst 2   : %q\n", s[:second+second])

		// Slice and print the last two runes of the strings
		// fmt.Printf("%s", strings[:])
		_, last := utf8.DecodeLastRuneInString(s)
		_, last2:= utf8.DecodeLastRuneInString(s[:len(s)-last])
		fmt.Printf("\tlast 2    : %q\n", s[len(s)-last-last2:])

		// Convert the string to []rune
		// Print the first and last two runes
		sRune := []rune(s)
		fmt.Printf("\tfirst 2   : %q\n", string(sRune[:2]))
		fmt.Printf("\tlast 2    : %q\n", string(sRune[len(sRune)-2:]))
	}
}
