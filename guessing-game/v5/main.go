package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	var input int
	for {
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			continue
		}

		guess := input
		if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again.")
		} else if guess > secretNumber {
			fmt.Println("Your guess is larger than the secret number. Please try again.")
		} else {
			fmt.Println("Congratulations! You guessed the secret number.")
			break
		}
	}
}
