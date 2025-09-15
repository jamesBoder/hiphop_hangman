package main

import (
	"fmt"
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

func main() {

	// Example word and attempts
	word := "GOLANG"
	guesses := 6

	// Setup initial state
	currentState := initWordState(word)

	// Game introduction
	fmt.Println("Welcome to HipHop Hangman!")
	fmt.Println("Guess the artist's name:")

	// Initial display
	displayWordState(currentState, guesses)

}
