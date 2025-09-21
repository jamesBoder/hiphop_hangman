package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// CLI Hiphop Hangman game

// Init word state. This function creates a slice of the same length as the word and fills it with underscores,
// but reveals spaces and hyphens as is.
func initWordState(word string) []string {
	state := make([]string, len(word))
	for i, char := range word {
		if char == ' ' || char == '-' {
			state[i] = string(char)
		} else {
			state[i] = "_"
		}
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

// create a random word function using names.go and math/rand
func randomWord(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	words := strings.Split(string(data), "\n")

	return words[rand.Intn(len(words))], nil
}

// Helper function to display guessed letters
func displayGuessedLetters(guessed []string) {
	fmt.Println("Guessed letters:", strings.Join(guessed, " "))
}

func main() {

	// Game Variables
	word, err := randomWord("names.txt")
	if err != nil {
		fmt.Println("Error reading word file:", err)
		return
	}

	// Convert word to uppercase for consistent comparison
	word = strings.ToUpper(word)

	guesses := 6
	playerScore := 0

	// Setup initial state
	currentState := initWordState(word)

	// Slice to track guessed letters
	var guessedLetters []string

	// Game introduction
	fmt.Println("Welcome to HipHop Hangman!")
	fmt.Println("Guess the artist's name:")

	// Initial display
	displayWordState(currentState, guesses)
	displayGuessedLetters(guessedLetters)

	// take user input using bufio and os packages
	scanner := bufio.NewScanner(os.Stdin)

	// receive user input
	for guesses > 0 {
		fmt.Print("Enter your guess (single letter or full word): ")
		scanner.Scan()
		guess := strings.ToUpper(scanner.Text())

		// Check if letter was already guessed
		if len(guess) == 1 {
			alreadyGuessed := false
			for _, g := range guessedLetters {
				if g == guess {
					alreadyGuessed = true
					break
				}
			}
			if alreadyGuessed {
				fmt.Println("You already guessed that letter. Try again.")
				continue
			}
			// Add to guessed letters
			guessedLetters = append(guessedLetters, guess)
		}

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
			displayGuessedLetters(guessedLetters)
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
