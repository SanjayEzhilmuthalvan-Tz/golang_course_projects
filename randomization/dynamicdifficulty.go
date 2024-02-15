package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game!

The program will pick some random numbers based on the number you give.
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
		fmt.Println("please enter a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please enter a positive number.")
		return
	}

	var balancer int
	if guess > 10 {
		balancer = guess / 4
	}

	for turn := maxTurns + balancer; turn > 0; turn-- {
		n := rand.Intn(guess) + 1

		if n == guess {
			fmt.Println("congratulations you won")
			return
		}
	}

	fmt.Println("you lost.. Wanna Try again?")
}
