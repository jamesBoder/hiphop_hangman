package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CLI Hiphop Hangman game

// Init word state. This function creates a slice of the same length as the word and fills it with underscores.
func initWordState(word string) []string {
	state := make([]string, len(word))
	for i := range state {
		state[i] = "_"
	}
	return state
}

// function to display current state of the word
func displayWordState(currentState []string, attempts int) {
	fmt.Println("Current word state:", strings.Join(currentState, " "))
	fmt.Println("Attempts left:", attempts)
}

// Create function to display hangman ASCII art based on remaining attempts
func displayHangman(attempts int) {
	stages := []string{
		`
		-----
		|   |
		|
		|
		|
		|
		=========
		`,
		`
		-----
		|   |
		|   O
		|
		|
		|
		=========
		`,
		`
		-----
		|   O
		|   |
		|   |
		|
		|
		=========
		`,
		`
		-----
		|   |
		|   O
		|  /|
		|
		|
		=========
		`,
		`
		-----
		|   |
		|   O
		|  /|\
		|  /
		|
		=========
		`,
		`
		-----
		|   |
		|   O
		|  /|\
		|  / \
		|
		=========
		`,
	}
	if attempts < len(stages) {
		fmt.Println(stages[len(stages)-1-attempts])
	}
}

func main() {

	// Game Variables
	word := "GOLANG"
	guesses := 6
	playerScore := 0

	// Setup initial state
	currentState := initWordState(word)

	// Game introduction
	fmt.Println("Welcome to HipHop Hangman!")
	fmt.Println("Guess the artist's name:")

	// Initial display
	displayWordState(currentState, guesses)

	// take user input using bufio and os packages
	scanner := bufio.NewScanner(os.Stdin)

	// receive user input
	for guesses > 0 {
		fmt.Print("Enter your guess (single letter or full word): ")
		scanner.Scan()
		guess := strings.ToUpper(scanner.Text())

		if len(guess) == 1 {
			// Single letter guess
			correctGuess := false
			for i, char := range word {
				if string(char) == guess {
					currentState[i] = guess
					correctGuess = true
				}
			}
			if !correctGuess {
				// deduct an attempt for wrong letter guess
				guesses--
				// Display hangman ASCII art
				displayHangman(guesses)
				fmt.Println("Wrong guess!")
			} else {
				fmt.Println("Good guess!")
			}

			// Show updated state of word after guess
			displayWordState(currentState, guesses)
		} else if len(guess) == len(word) {
			// Full word guess
			if guess == word {
				// Add score for correct full word guess
				playerScore += 10
				fmt.Println("Congratulations! You've guessed the word:", word)
				break
			} else {
				// deduct an attempt for wrong full word guess
				guesses--
				// Display hangman ASCII art
				displayHangman(guesses)
				fmt.Println("Wrong guess!")
			}
		} else {
			fmt.Println("Invalid input. Please enter a single letter or the full word.")
			continue
		}

		// Check for win condition
		if strings.Join(currentState, "") == word {
			fmt.Println("Congratulations! You've guessed the word:", word)
			fmt.Println("Your score is:", playerScore)
			break
		}
	}

	if guesses == 0 {
		fmt.Println("Game over! The correct word was:", word)
	}

}
