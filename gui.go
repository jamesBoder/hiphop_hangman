package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type GameGUI struct {
	app          fyne.App
	window       fyne.Window
	game         *GameState
	selectedFile string
	themeManager *ThemeManager

	// UI Elements
	titleLabel      *widget.Label
	wordLabel       *widget.Label
	hangmanLabel    *widget.Label
	attemptsLabel   *widget.Label
	guessedLabel    *widget.Label
	scoreLabel      *widget.Label
	guessEntry      *widget.Entry
	guessButton     *widget.Button
	newGameButton   *widget.Button
	categorySelect  *widget.Select
	themeSelect     *widget.Select
	switchGUIButton *widget.Button

	// Containers
	gameContainer     *fyne.Container
	categoryContainer *fyne.Container
	mainContainer     *fyne.Container
}

func createGUI() {
	gui := &GameGUI{}
	gui.app = app.New()
	gui.window = gui.app.NewWindow("🎤 Hip-Hop Hangman 🎤")
	gui.window.Resize(fyne.NewSize(900, 700))

	gui.setupUI()
	gui.showCategorySelection()

	gui.window.ShowAndRun()
}

func (gui *GameGUI) setupUI() {
	// Create UI elements
	gui.titleLabel = widget.NewLabel("🎤 Hip-Hop Hangman - Basic GUI 🎤")
	gui.titleLabel.Alignment = fyne.TextAlignCenter
	gui.titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	gui.wordLabel = widget.NewLabel("")
	gui.wordLabel.Alignment = fyne.TextAlignCenter
	gui.wordLabel.TextStyle = fyne.TextStyle{Monospace: true}

	gui.hangmanLabel = widget.NewLabel("")
	gui.hangmanLabel.Alignment = fyne.TextAlignCenter
	gui.hangmanLabel.TextStyle = fyne.TextStyle{Monospace: true}

	gui.attemptsLabel = widget.NewLabel("")
	gui.attemptsLabel.Alignment = fyne.TextAlignCenter

	gui.guessedLabel = widget.NewLabel("")
	gui.guessedLabel.Alignment = fyne.TextAlignCenter

	gui.scoreLabel = widget.NewLabel("Score: 0")
	gui.scoreLabel.Alignment = fyne.TextAlignCenter

	// Create a single-line entry field that responds to Enter key
	gui.guessEntry = widget.NewEntry()
	gui.guessEntry.SetPlaceHolder("🎤 Enter a letter or full word here... 🎤")
	gui.guessEntry.OnSubmitted = func(text string) {
		gui.makeGuess()
	}
	// Make text larger and bold for better visibility
	gui.guessEntry.TextStyle = fyne.TextStyle{Bold: true}

	gui.guessButton = widget.NewButton("Make Guess", gui.makeGuess)
	gui.newGameButton = widget.NewButton("New Game", gui.showCategorySelection)

	// Theme selector
	if gui.themeManager != nil {
		gui.themeSelect = widget.NewSelect(gui.themeManager.GetThemeNames(), func(selected string) {
			gui.themeManager.SetTheme(selected)
		})
		gui.themeSelect.SetSelected("Hip-Hop")
	}

	// GUI style switcher
	gui.switchGUIButton = widget.NewButton("🎨 Switch to Modern GUI", func() {
		manager := GetGUIManager()
		// Update game state in manager before switching
		if gui.game != nil {
			manager.UpdateGameState(gui.game, gui.selectedFile)
		}
		manager.SwitchToModernGUI()
	})

	// Category selection
	categoryOptions := make([]string, 0, len(CategoryNames))
	for i := 1; i <= 7; i++ {
		key := fmt.Sprintf("%d", i)
		name := CategoryNames[key]
		file := Categories[key]
		count := GetCategoryStats(file)
		option := fmt.Sprintf("%s: %s (%d artists)", key, name, count)
		categoryOptions = append(categoryOptions, option)
	}

	gui.categorySelect = widget.NewSelect(categoryOptions, gui.onCategorySelected)
	gui.categorySelect.PlaceHolder = "Choose a category..."
}

func (gui *GameGUI) showCategorySelection() {
	welcomeLabel := widget.NewLabel("Welcome to the ultimate hip-hop artist guessing game!")
	welcomeLabel.Alignment = fyne.TextAlignCenter

	instructionLabel := widget.NewLabel("Select a category to start playing:")
	instructionLabel.Alignment = fyne.TextAlignCenter

	// Create controls container
	controlsContainer := container.NewVBox()

	// Add theme selector if available
	if gui.themeSelect != nil {
		themeLabel := widget.NewLabel("🎨 Choose Theme:")
		themeLabel.Alignment = fyne.TextAlignCenter
		controlsContainer.Add(themeLabel)
		controlsContainer.Add(gui.themeSelect)
		controlsContainer.Add(widget.NewSeparator())
	}

	// Add GUI switcher
	guiLabel := widget.NewLabel("🔄 GUI Style:")
	guiLabel.Alignment = fyne.TextAlignCenter
	controlsContainer.Add(guiLabel)
	controlsContainer.Add(gui.switchGUIButton)

	gui.categoryContainer = container.NewVBox(
		gui.titleLabel,
		widget.NewSeparator(),
		welcomeLabel,
		widget.NewSeparator(),
		instructionLabel,
		gui.categorySelect,
		widget.NewSeparator(),
		controlsContainer,
	)

	gui.window.SetContent(gui.categoryContainer)
}

func (gui *GameGUI) onCategorySelected(selected string) {
	if selected == "" {
		return
	}

	// Extract category number from selection
	parts := strings.Split(selected, ":")
	if len(parts) == 0 {
		return
	}

	categoryKey := strings.TrimSpace(parts[0])
	if file, exists := Categories[categoryKey]; exists {
		gui.selectedFile = file
		gui.startNewGame()
	}
}

func (gui *GameGUI) startNewGame() {
	if gui.selectedFile == "" {
		return
	}

	// Get random word from selected category
	word, err := RandomWord(gui.selectedFile)
	if err != nil {
		dialog.ShowError(err, gui.window)
		return
	}

	// Create new game
	gui.game = NewGame(word)

	// Setup game UI
	gui.setupGameUI()
	gui.updateGameDisplay()
}

func (gui *GameGUI) setupGameUI() {
	// Enable the input field
	gui.guessEntry.Enable()

	// Left side - Hangman art (centered and equal space)
	hangmanTitle := widget.NewLabel("🎭 Hangman Status:")
	hangmanTitle.Alignment = fyne.TextAlignCenter
	hangmanTitle.TextStyle = fyne.TextStyle{Bold: true}

	gui.hangmanLabel.Resize(fyne.NewSize(400, 350))
	leftSide := container.NewVBox(
		hangmanTitle,
		widget.NewSeparator(),
		container.NewCenter(gui.hangmanLabel),
	)

	// Right side - Word and game info (centered and equal space)
	wordTitle := widget.NewLabel("🎯 Word to Guess:")
	wordTitle.Alignment = fyne.TextAlignCenter
	wordTitle.TextStyle = fyne.TextStyle{Bold: true}

	gui.wordLabel.Resize(fyne.NewSize(400, 80))
	gui.wordLabel.TextStyle = fyne.TextStyle{Bold: true}

	rightSide := container.NewVBox(
		wordTitle,
		widget.NewSeparator(),
		container.NewCenter(gui.wordLabel),
		widget.NewSeparator(),
		container.NewCenter(gui.scoreLabel),
		container.NewCenter(gui.attemptsLabel),
		widget.NewSeparator(),
		container.NewCenter(gui.guessedLabel),
	)

	// Main game area with centered equal-width containers using GridWithColumns
	gameArea := container.NewGridWithColumns(2,
		container.NewCenter(leftSide),
		container.NewCenter(rightSide),
	)

	// Input container with centered elements and larger button
	inputLabel := widget.NewLabel("💭 Enter your guess (letter or full word):")
	inputLabel.Alignment = fyne.TextAlignCenter
	inputLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Create a simple, properly centered input field
	// Use VBox with proper spacing and centering

	// Make the guess button larger and more prominent
	gui.guessButton.Resize(fyne.NewSize(200, 50))

	// Use a grid layout to force the input field to take up more space
	// Create empty spacers and put the input field in the middle column
	leftSpacer := widget.NewLabel("")
	rightSpacer := widget.NewLabel("")

	// Use 3-column grid: 20% - 60% - 20% to force input field to be wider
	inputFieldRow := container.NewGridWithColumns(3,
		leftSpacer,     // 33% width
		gui.guessEntry, // 33% width - this should force expansion
		rightSpacer,    // 33% width
	)

	inputContainer := container.NewVBox(
		container.NewCenter(inputLabel),
		widget.NewSeparator(),
		inputFieldRow,
		widget.NewSeparator(),
		container.NewCenter(gui.guessButton),
	)

	// Button container with game controls and GUI switcher
	gameButtonsRow := container.NewGridWithColumns(2, gui.newGameButton, gui.switchGUIButton)

	// Theme selector row if available
	var themeRow *fyne.Container
	if gui.themeSelect != nil {
		themeLabel := widget.NewLabel("🎨 Theme:")
		themeLabel.Alignment = fyne.TextAlignCenter
		themeRow = container.NewVBox(
			themeLabel,
			gui.themeSelect,
		)
	}

	// Button container with all controls
	buttonContainer := container.NewVBox(
		gameButtonsRow,
	)

	if themeRow != nil {
		buttonContainer.Add(widget.NewSeparator())
		buttonContainer.Add(themeRow)
	}

	// Main game container with everything centered
	gui.gameContainer = container.NewVBox(
		container.NewCenter(gui.titleLabel),
		widget.NewSeparator(),
		container.NewCenter(gameArea),
		widget.NewSeparator(),
		inputContainer,
		widget.NewSeparator(),
		container.NewCenter(buttonContainer),
	)

	gui.window.SetContent(gui.gameContainer)
}

func (gui *GameGUI) updateGameDisplay() {
	if gui.game == nil {
		return
	}

	// Update word display
	gui.wordLabel.SetText(gui.game.GetWordDisplay())

	// Update hangman art with emojis for better display
	gui.hangmanLabel.SetText(GetEmojiHangmanArt(gui.game.Attempts))

	// Update attempts
	gui.attemptsLabel.SetText(fmt.Sprintf("❤️ Lives remaining: %d", gui.game.Attempts))

	// Update guessed letters
	if len(gui.game.GuessedLetters) > 0 {
		gui.guessedLabel.SetText(fmt.Sprintf("🔤 Guessed letters: %s", gui.game.GetGuessedLettersDisplay()))
	} else {
		gui.guessedLabel.SetText("🔤 Guessed letters: none")
	}

	// Update score
	gui.scoreLabel.SetText(fmt.Sprintf("🏆 Score: %d", gui.game.Score))

	// Check game over conditions
	if gui.game.IsGameOver {
		gui.handleGameOver()
	}
}

func (gui *GameGUI) makeGuess() {
	if gui.game == nil || gui.game.IsGameOver {
		return
	}

	guess := strings.TrimSpace(gui.guessEntry.Text)
	if guess == "" {
		gui.showCustomDialog("Invalid Input", "Please enter a letter or word to guess.\n\nPress Enter or click OK to continue.")
		return
	}

	// Check if single letter was already guessed
	if len(guess) == 1 {
		for _, g := range gui.game.GuessedLetters {
			if strings.ToUpper(g) == strings.ToUpper(guess) {
				gui.showCustomDialog("Already Guessed", "You already guessed that letter. Try again.\n\nPress Enter or click OK to continue.")
				gui.guessEntry.SetText("")
				return
			}
		}
	}

	// Make the guess
	correctGuess := gui.game.MakeGuess(guess)

	// Clear the entry
	gui.guessEntry.SetText("")

	// Show feedback
	if len(guess) == 1 {
		if correctGuess {
			// Don't show dialog for correct single letter guesses to keep game flowing
		} else {
			// Show brief feedback for wrong guesses with Enter key support
			go func() {
				gui.showCustomDialog("Wrong Guess", "That letter is not in the word.\n\nPress Enter or click OK to continue.")
			}()
		}
	} else {
		// Full word guess feedback
		if correctGuess {
			// Will be handled in updateGameDisplay -> handleGameOver
		} else {
			go func() {
				gui.showCustomDialog("Wrong Guess", "That's not the correct word.\n\nPress Enter or click OK to continue.")
			}()
		}
	}

	// Update display
	gui.updateGameDisplay()
}

// showCustomDialog creates a custom dialog that responds to Enter key
func (gui *GameGUI) showCustomDialog(title, message string) {
	// Create a label for the message
	messageLabel := widget.NewLabel(message)
	messageLabel.Alignment = fyne.TextAlignCenter
	messageLabel.Wrapping = fyne.TextWrapWord

	// Create OK button
	okButton := widget.NewButton("OK", nil)

	// Create the dialog content
	content := container.NewVBox(
		messageLabel,
		widget.NewSeparator(),
		container.NewCenter(okButton),
	)

	// Create the dialog
	d := dialog.NewCustom(title, "", content, gui.window)
	d.Resize(fyne.NewSize(400, 200))

	// Set up the OK button to close dialog
	okButton.OnTapped = func() {
		d.Hide()
	}

	// Add keyboard shortcut for Enter key
	gui.window.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if key.Name == fyne.KeyReturn || key.Name == fyne.KeyEnter {
			d.Hide()
			// Reset the keyboard handler after dialog closes
			gui.window.Canvas().SetOnTypedKey(nil)
		}
	})

	// Show the dialog
	d.Show()
}

func (gui *GameGUI) handleGameOver() {
	// Disable input
	gui.guessEntry.Disable()
	gui.guessButton.Disable()

	// Show game over message
	var message string
	var title string

	if gui.game.IsWon {
		title = "🎉 Congratulations! 🎉"
		message = fmt.Sprintf("You guessed the word: %s\nYour score: %d points", gui.game.Word, gui.game.Score)
	} else {
		title = "💀 Game Over 💀"
		message = fmt.Sprintf("The correct word was: %s\nBetter luck next time!", gui.game.Word)
	}

	// Show dialog with option to play again
	dialog.ShowConfirm(title, message+"\n\nWould you like to play again?",
		func(playAgain bool) {
			if playAgain {
				gui.showCategorySelection()
			} else {
				gui.app.Quit()
			}
		}, gui.window)
}
