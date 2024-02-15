package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 10
	usage    = `Welcome to the Lucky Number Game!

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("you haven't entered a number pleasae enter a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please enter a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if n == guess {
			switch rand.Intn(4) {
			case 0:
				fmt.Println("you are the winner")
			case 1:
				fmt.Println("you are lucky")
			case 2:
				fmt.Println("you are awesome")
			case 3:
				fmt.Println("you are perfect")
			}
			return
		}
	}

	msg := "%s Try again?\n"

	switch rand.Intn(3) {
	case 0:
		fmt.Printf(msg, "you lost ")
	case 1:
		fmt.Printf(msg, "Better luck next time")
	case 2:
		fmt.Println("you may try again")
	}

}
