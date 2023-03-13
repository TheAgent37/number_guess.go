package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	if !play_again {
		fmt.Printf(`
Welcome to my number guessing game!

Rules:

-You have %d tries
-Random generated number is between %d and %d (both including)
-You can play the game how many as you want!

`, MAX_GUESS, MIN_NUM, MAX_NUM)
	}
	guess()
	try_again(win)
}

const MAX_GUESS = 3
const MIN_NUM = 1
const MAX_NUM = 3

var win bool = false
var play_again bool = false

func guess() {
	win = false
	var guess_count int
	num := rand.Intn(MAX_NUM-MIN_NUM+1) + MIN_NUM
	fmt.Println("Random number created between", MIN_NUM, "and", MAX_NUM)
	for guess_count < MAX_GUESS && !win {
		gnum := ""
		fmt.Printf("You have %d tries left. Enter your guess: ", MAX_GUESS-guess_count)
		fmt.Scan(&gnum)
		n, err := strconv.Atoi(gnum)
		if err != nil {
			fmt.Printf("Error! Please enter a number.\n\n")
			continue
		}
		guess_count += 1

		switch {
		case n > num:
			fmt.Printf("Your guess is too high.\n\n")
		case n < num:
			fmt.Printf("Your guess is too low.\n\n")
		default:
			fmt.Printf("You guessed the number (%d) in %d tries!\n\n", num, guess_count)
			win = true
		}
	}
	if win == false {
		fmt.Println("You couldn't guess the number!")
	}
}

func try_again(win bool) {
	var try_loop bool = true
	for try_loop {
		var play string
		if win {
			fmt.Printf("Congratulations! Do you want to play again? (y/n): ")
		} else {
			fmt.Printf("Do you want to try again? (y/n): ")
		}
		fmt.Scan(&play)
		switch play {
		case "y":
			play_again = true
			main()
		case "n":
			fmt.Println("Goodbye!")
			try_loop = false
			break
		default:
			fmt.Printf("Invalid input.\n\n")
		}
	}
}
