package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Game state structure
type GameState struct {
	Word           string
	CurrentState   []string
	GuessedLetters []string
	Attempts       int
	Score          int
	IsGameOver     bool
	IsWon          bool
}

// Category mapping
var Categories = map[string]string{
	"1": "east_coast.txt",
	"2": "west_coast.txt",
	"3": "south.txt",
	"4": "midwest.txt",
	"5": "international.txt",
	"6": "groups.txt",
	"7": "names.txt", // All artists
}

var CategoryNames = map[string]string{
	"1": "East Coast (NY, NJ, PA, MD, VA, CT)",
	"2": "West Coast (CA, WA, OR, NV)",
	"3": "South (GA, TX, FL, AL, MS, LA, NC, SC, TN)",
	"4": "Midwest (IL, MI, OH, MN, WI, IN)",
	"5": "International (Non-US artists)",
	"6": "Groups (Collective groups)",
	"7": "All Artists (Complete collection)",
}

// GetCategoryStats returns the number of artists in a category file
func GetCategoryStats(file string) int {
	data, err := os.ReadFile(file)
	if err != nil {
		return 0
	}
	words := strings.Split(strings.TrimSpace(string(data)), "\n")
	return len(words)
}

// DisplayCategories shows available categories with artist counts (for CLI)
func DisplayCategories() {
	fmt.Println("=== Welcome to Hiphop Hangman ===")
	fmt.Println("\n=== Choose Your Hip Hop Category ===")
	for key, name := range CategoryNames {
		file := Categories[key]
		count := GetCategoryStats(file)
		fmt.Printf("%s: %s (%d artists)\n", key, name, count)
	}
	fmt.Println()
}

// SelectCategory prompts user to choose a category and returns the file path (for CLI)
func SelectCategory(scanner *bufio.Scanner) string {
	for {
		fmt.Print("Enter category number (1-7): ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		if file, exists := Categories[choice]; exists {
			categoryName := CategoryNames[choice]
			fmt.Printf("Selected: %s\n\n", categoryName)
			return file
		}
		fmt.Println("Invalid choice. Please enter a number between 1 and 7.")
	}
}

// InitWordState creates a slice of the same length as the word and fills it with underscores,
// but reveals spaces, hyphens, and periods as is.
func InitWordState(word string) []string {
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

// DisplayWordState displays current state of the word (for CLI)
func DisplayWordState(currentState []string, attempts int) {
	fmt.Println("Current word state:", strings.Join(currentState, " "))
	fmt.Println("Attempts left:", attempts)
}

// GetHangmanArt returns hangman ASCII art based on remaining attempts
func GetHangmanArt(attempts int) string {
	stages := []string{
		`     ┌─────────┐
    │         │
    │         
    │         
    │         
    │         
    │         
    │         
 ───┴───      `,
		`     ┌─────────┐
    │         │
    │         ●
    │         
    │         
    │         
    │         
    │         
 ───┴───      `,
		`     ┌─────────┐
    │         │
    │         ●
    │         │
    │         │
    │         
    │         
    │         
 ───┴───      `,
		`     ┌─────────┐
    │         │
    │         ●
    │        ╱│
    │         │
    │         
    │         
    │         
 ───┴───      `,
		`     ┌─────────┐
    │         │
    │         ●
    │        ╱│╲
    │         │
    │         
    │         
    │         
 ───┴───      `,
		`     ┌─────────┐
    │         │
    │         ●
    │        ╱│╲
    │         │
    │        ╱ 
    │         
    │         
 ───┴───      `,
		`     ┌─────────┐
    │         │
    │         ●
    │        ╱│╲
    │         │
    │        ╱ ╲
    │         
    │         
 ───┴───      `,
	}

	// Handle edge cases
	if attempts < 0 {
		return stages[len(stages)-1] // Show complete hangman for negative attempts
	}
	if attempts >= len(stages) {
		return stages[0] // Show empty gallows for high attempts
	}

	return stages[len(stages)-1-attempts]
}

// GetEmojiHangmanArt returns emoji-based hangman art for better GUI display
func GetEmojiHangmanArt(attempts int) string {
	stages := []string{
		// Stage 0 - Empty gallows (6 attempts remaining)
		`🏗️ Gallows Ready 🏗️
		
🎯 Let's Play! 🎯
		
🎤 Guess the Hip-Hop Artist! 🎤`,

		// Stage 1 - Head (5 attempts remaining)
		`🏗️ Gallows 🏗️
		
😵 Uh oh!
		
💀 Getting serious...`,

		// Stage 2 - Body (4 attempts remaining)
		`🏗️ Gallows 🏗️
		
😵 Head
🦴 Body
		
⚠️ Danger zone!`,

		// Stage 3 - Left arm (3 attempts remaining)
		`🏗️ Gallows 🏗️
		
😵 Head
💪🦴 Body + Arm
		
🚨 Critical!`,

		// Stage 4 - Both arms (2 attempts remaining)
		`🏗️ Gallows 🏗️
		
😵 Head
💪🦴💪 Full torso
		
🔥 Last chances!`,

		// Stage 5 - Left leg (1 attempt remaining)
		`🏗️ Gallows 🏗️
		
😵 Head
💪🦴💪 Torso
🦵 One leg
		
💥 FINAL WARNING!`,

		// Stage 6 - Complete hangman (0 attempts - game over)
		`🏗️ Gallows 🏗️
		
💀 RIP
💪🦴💪 Complete
🦵 🦵 Both legs
		
☠️ GAME OVER! ☠️`,
	}

	// Handle edge cases
	if attempts < 0 {
		return stages[len(stages)-1] // Show complete hangman for negative attempts
	}
	if attempts >= len(stages) {
		return stages[0] // Show empty gallows for high attempts
	}

	return stages[len(stages)-1-attempts]
}

// DisplayHangman displays hangman ASCII art (for CLI)
func DisplayHangman(attempts int) {
	fmt.Println(GetHangmanArt(attempts))
}

// RandomWord returns a random word from the specified file
func RandomWord(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	words := strings.Split(string(data), "\n")

	// Filter out empty lines
	var validWords []string
	for _, word := range words {
		trimmed := strings.TrimSpace(word)
		if trimmed != "" {
			validWords = append(validWords, trimmed)
		}
	}

	if len(validWords) == 0 {
		return "", fmt.Errorf("no words found in file")
	}

	return validWords[rand.Intn(len(validWords))], nil
}

// DisplayGuessedLetters displays guessed letters (for CLI)
func DisplayGuessedLetters(guessed []string) {
	fmt.Println("Guessed letters:", strings.Join(guessed, " "))
}

// NewGame creates a new game state with the given word
func NewGame(word string) *GameState {
	return &GameState{
		Word:           strings.ToUpper(word),
		CurrentState:   InitWordState(strings.ToUpper(word)),
		GuessedLetters: []string{},
		Attempts:       6,
		Score:          0,
		IsGameOver:     false,
		IsWon:          false,
	}
}

// MakeGuess processes a guess and updates the game state
func (g *GameState) MakeGuess(guess string) bool {
	if g.IsGameOver {
		return false
	}

	guess = strings.ToUpper(guess)

	if len(guess) == 1 {
		// Single letter guess
		// Check if already guessed
		for _, g := range g.GuessedLetters {
			if g == guess {
				return false // Already guessed
			}
		}

		g.GuessedLetters = append(g.GuessedLetters, guess)
		correctGuess := false

		for i, char := range g.Word {
			if string(char) == guess {
				g.CurrentState[i] = guess
				correctGuess = true
			}
		}

		if !correctGuess {
			g.Attempts--
		}

		// Check win condition
		if strings.Join(g.CurrentState, "") == g.Word {
			g.IsWon = true
			g.IsGameOver = true
			g.Score += 10
		}

		// Check lose condition
		if g.Attempts <= 0 {
			g.IsGameOver = true
		}

		return correctGuess
	} else if len(guess) == len(g.Word) {
		// Full word guess
		if guess == g.Word {
			g.IsWon = true
			g.IsGameOver = true
			g.Score += 10
			// Fill in the word
			for i, char := range g.Word {
				g.CurrentState[i] = string(char)
			}
			return true
		} else {
			g.Attempts--
			if g.Attempts <= 0 {
				g.IsGameOver = true
			}
			return false
		}
	}

	return false
}

// GetWordDisplay returns the current word state as a string
func (g *GameState) GetWordDisplay() string {
	return strings.Join(g.CurrentState, " ")
}

// GetGuessedLettersDisplay returns guessed letters as a string
func (g *GameState) GetGuessedLettersDisplay() string {
	return strings.Join(g.GuessedLetters, " ")
}

// RunCLIGame runs the CLI version of the game
func RunCLIGame() {
	// take user input using bufio and os packages
	scanner := bufio.NewScanner(os.Stdin)

	// Category selection
	DisplayCategories()
	selectedFile := SelectCategory(scanner)

	// Game Variables
	word, err := RandomWord(selectedFile)
	if err != nil {
		fmt.Println("Error reading word file:", err)
		return
	}

	// Create new game
	game := NewGame(word)

	// Game introduction
	fmt.Println("Guess the artist's name:")

	// Initial display
	DisplayWordState(game.CurrentState, game.Attempts)
	DisplayGuessedLetters(game.GuessedLetters)

	// receive user input
	for !game.IsGameOver {
		fmt.Print("Enter your guess (single letter or full word): ")
		scanner.Scan()
		guess := strings.TrimSpace(scanner.Text())

		if guess == "" {
			fmt.Println("Please enter a valid guess.")
			continue
		}

		// Check if letter was already guessed
		if len(guess) == 1 {
			alreadyGuessed := false
			for _, g := range game.GuessedLetters {
				if strings.ToUpper(g) == strings.ToUpper(guess) {
					alreadyGuessed = true
					break
				}
			}
			if alreadyGuessed {
				fmt.Println("You already guessed that letter. Try again.")
				continue
			}
		}

		correctGuess := game.MakeGuess(guess)

		if len(guess) == 1 {
			if !correctGuess {
				DisplayHangman(game.Attempts)
				fmt.Println("Wrong guess!")
			} else {
				fmt.Println("Good guess!")
			}
		} else {
			if !correctGuess {
				DisplayHangman(game.Attempts)
				fmt.Println("Wrong guess!")
			}
		}

		// Show updated state of word after guess
		if !game.IsGameOver {
			DisplayWordState(game.CurrentState, game.Attempts)
			DisplayGuessedLetters(game.GuessedLetters)
		}
	}

	if game.IsWon {
		fmt.Println("Congratulations! You've guessed the word:", game.Word)
		fmt.Println("Your score is:", game.Score)
	} else {
		fmt.Println("Game over! The correct word was:", game.Word)
	}
}

func main() {
	// Start the GUI version by default
	createGUI()
}
