package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1
	var guess int
	fmt.Println("Number Guessing Game.")
	fmt.Println("Guess a number between 100 and 1.")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		if guess < secretNumber {
			fmt.Println("Too low")
		} else if guess > secretNumber {
			fmt.Println("Too high")
		} else {
			fmt.Println("Correct! you won")
		}
	}
}
