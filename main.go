package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// CLI Hiphop Hangman game

// Category mapping
var categories = map[string]string{
	"1": "east_coast.txt",
	"2": "west_coast.txt",
	"3": "south.txt",
	"4": "midwest.txt",
	"5": "international.txt",
	"6": "groups.txt",
	"7": "names.txt", // All artists
}

var categoryNames = map[string]string{
	"1": "East Coast (NY, NJ, PA, MD, VA, CT)",
	"2": "West Coast (CA, WA, OR, NV)",
	"3": "South (GA, TX, FL, AL, MS, LA, NC, SC, TN)",
	"4": "Midwest (IL, MI, OH, MN, WI, IN)",
	"5": "International (Non-US artists)",
	"6": "Groups (Collective groups)",
	"7": "All Artists (Complete collection)",
}

// getCategoryStats returns the number of artists in a category file
func getCategoryStats(file string) int {
	data, err := os.ReadFile(file)
	if err != nil {
		return 0
	}
	words := strings.Split(strings.TrimSpace(string(data)), "\n")
	return len(words)
}

// displayCategories shows available categories with artist counts
func displayCategories() {
	fmt.Println("=== Welcome to Hiphop Hangman ===")
	fmt.Println("\n=== Choose Your Hip Hop Category ===")
	for key, name := range categoryNames {
		file := categories[key]
		count := getCategoryStats(file)
		fmt.Printf("%s: %s (%d artists)\n", key, name, count)
	}
	fmt.Println()
}

// selectCategory prompts user to choose a category and returns the file path
func selectCategory(scanner *bufio.Scanner) string {
	for {
		fmt.Print("Enter category number (1-7): ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		if file, exists := categories[choice]; exists {
			categoryName := categoryNames[choice]
			fmt.Printf("Selected: %s\n\n", categoryName)
			return file
		}
		fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
	}
}

// Init word state. This function creates a slice of the same length as the word and fills it with underscores,
// but reveals spaces, hyphens, and periods as is.
func initWordState(word string) []string {
	state := make([]string, len(word))
	for i, char := range word {
		if char == ' ' || char == '-' || char == '.' {
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
	// take user input using bufio and os packages
	scanner := bufio.NewScanner(os.Stdin)

	// Category selection
	displayCategories()
	selectedFile := selectCategory(scanner)

	// Game Variables
	word, err := randomWord(selectedFile)
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
	// fmt.Println("Welcome to HipHop Hangman!")
	fmt.Println("Guess the artist's name:")

	// Initial display
	displayWordState(currentState, guesses)
	displayGuessedLetters(guessedLetters)

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
