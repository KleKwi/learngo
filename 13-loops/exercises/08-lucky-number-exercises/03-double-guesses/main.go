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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurn = 5
	usage   = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess two of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf(usage, maxTurn)
		return
	}

	guess1, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Wrong number.")
		return
	}

	var guess2 int
	if len(args) == 2 {
		guess2, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Wrong number.")
			return
		}
	}

	if guess1 < 0 || guess2 <= 0 {
		fmt.Println("Give me positive numbers.")
		return
	}

	max := guess1
	if guess1 > guess2 {
		max = guess2
	}

	for turn := 1; turn <= maxTurn; turn++ {
		n := rand.Intn(max + 1)
		if n == guess1 || n == guess2 {
			fmt.Println("YOU WIN")
			return
		}
	}
	fmt.Println("YOU LOST. TRY AGAIN?")
}
