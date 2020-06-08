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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
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
			fmt.Println("You Win!")
		}
		return
	}
	fmt.Println("You Lost ... Try again?")
}
