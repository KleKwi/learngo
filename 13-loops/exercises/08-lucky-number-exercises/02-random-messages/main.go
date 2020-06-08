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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

const (
	maxTurn = 5
	usage   = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 2 {
		fmt.Printf(usage, maxTurn)
		return
	}

	guess, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Wrong number.")
		return
	}

	for turn := 1; turn <= maxTurn; turn++ {
		n := rand.Intn(guess + 1)
		if n != guess {
			continue
		}
		if turn == 1 {
			fmt.Println("First Time Winner!")
		} else {
			switch n := rand.Intn(2); {
			case n == 0:
				fmt.Println("YOU WON")
			case n == 1:
				fmt.Println("YOU'RE AWESOME")
			}
		}
		return
	}
	switch n := rand.Intn(2); {
	case n == 0:
		fmt.Println("LOSER!")
	case n == 1:
		fmt.Println("YOU LOST. TRY AGAIN?")
	}
}
