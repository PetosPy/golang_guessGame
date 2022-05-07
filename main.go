package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get user input from terminal
	r := bufio.NewReader(os.Stdin)

	var userNum int

	// Prompt  the user to enter a number

	rand.Seed(time.Now().UnixNano())
	correctNum := rand.Intn(11)
	fmt.Println("Enter a number between 1 and 10:")

	for userNum != correctNum {

		// Read user input
		text, err := r.ReadString('\n')
		if err != nil {
			panic(err)

		}
		textToInt, err := strconv.Atoi(strings.TrimSpace(text))

		if err != nil {
			panic(err)
		}

		userNum = textToInt

		// Check if user obeyed the rules
		boundaryCheck := userNum >= 1 && userNum <= 10

		// Check if the user is right
		checkIfRight := userNum == correctNum

		if boundaryCheck {
			if checkIfRight {
				fmt.Println("You guessed it!")
				fmt.Printf("The number %v matched our number(%v)!\n", userNum, correctNum)
			} else {
				fmt.Printf("The number %v is wrong!: \n", textToInt)
			}
		} else {
			fmt.Printf("The number %v is out of bounds. Try again!! \n", textToInt)
		}

	}
}
