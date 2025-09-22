package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ModernGameGUI implements a modern Material Design-inspired interface
type ModernGameGUI struct {
	app          fyne.App
	window       fyne.Window
	game         *GameState
	selectedFile string
	themeManager *ThemeManager

	// UI Elements
	titleCard   *fyne.Container
	gameCard    *fyne.Container
	inputCard   *fyne.Container
	statsCard   *fyne.Container
	hangmanCard *fyne.Container

	// Widgets
	titleLabel      *widget.Label
	wordLabel       *widget.Label
	hangmanLabel    *widget.Label
	attemptsLabel   *widget.Label
	guessedLabel    *widget.Label
	scoreLabel      *widget.Label
	progressBar     *widget.ProgressBar
	guessEntry      *widget.Entry
	guessButton     *widget.Button
	newGameButton   *widget.Button
	themeSelect     *widget.Select
	categorySelect  *widget.Select
	switchGUIButton *widget.Button

	// Advanced widgets
	tabs    *container.AppTabs
	toolbar *widget.Toolbar

	// Containers
	mainContainer *fyne.Container
	sidePanel     *fyne.Container
}

func createModernGUI() {
	gui := &ModernGameGUI{}
	gui.app = app.New()
	gui.app.SetIcon(nil) // Could set a custom icon here

	// Initialize theme manager
	gui.themeManager = NewThemeManager(gui.app)
	gui.themeManager.SetTheme("Hip-Hop") // Start with custom theme

	gui.window = gui.app.NewWindow("🎤 Hip-Hop Hangman - Modern Edition 🎤")
	gui.window.Resize(fyne.NewSize(1200, 800))
	gui.window.SetMaster() // Makes this the main window

	gui.setupModernUI()
	gui.showModernCategorySelection()

	gui.window.ShowAndRun()
}

func (gui *ModernGameGUI) setupModernUI() {
	// Create toolbar
	gui.createToolbar()

	// Create cards (Material Design inspired containers)
	gui.createCards()

	// Create advanced widgets
	gui.createAdvancedWidgets()

	// Setup theme selector
	gui.setupThemeSelector()
}

func (gui *ModernGameGUI) createToolbar() {
	gui.toolbar = widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			gui.showModernCategorySelection()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			gui.showSettingsDialog()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), func() {
			gui.showAboutDialog()
		}),
	)
}

func (gui *ModernGameGUI) createCards() {
	// Title Card - Elevated header
	gui.titleLabel = widget.NewLabel("🎤 Hip-Hop Hangman - Modern GUI 🎤")
	gui.titleLabel.Alignment = fyne.TextAlignCenter
	gui.titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	gui.titleCard = container.NewVBox(
		gui.titleLabel,
		widget.NewSeparator(),
	)

	// Stats Card - Game statistics
	gui.scoreLabel = widget.NewLabel("Score: 0")
	gui.attemptsLabel = widget.NewLabel("Lives: 6")
	gui.progressBar = widget.NewProgressBar()
	gui.progressBar.SetValue(1.0) // Start at full

	gui.statsCard = container.NewVBox(
		widget.NewCard("📊 Game Stats", "", container.NewVBox(
			gui.scoreLabel,
			gui.attemptsLabel,
			widget.NewLabel("Progress:"),
			gui.progressBar,
		)),
	)

	// Hangman Card - Visual game state
	gui.hangmanLabel = widget.NewLabel("")
	gui.hangmanLabel.Alignment = fyne.TextAlignCenter
	gui.hangmanLabel.TextStyle = fyne.TextStyle{Monospace: true}

	gui.hangmanCard = container.NewVBox(
		widget.NewCard("🎭 Hangman", "", container.NewCenter(gui.hangmanLabel)),
	)

	// Game Card - Main game area
	gui.wordLabel = widget.NewLabel("")
	gui.wordLabel.Alignment = fyne.TextAlignCenter
	gui.wordLabel.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}

	gui.guessedLabel = widget.NewLabel("")
	gui.guessedLabel.Alignment = fyne.TextAlignCenter

	gui.gameCard = container.NewVBox(
		widget.NewCard("🎯 Word to Guess", "", container.NewVBox(
			container.NewCenter(gui.wordLabel),
			widget.NewSeparator(),
			gui.guessedLabel,
		)),
	)

	// Input Card - User interaction
	gui.guessEntry = widget.NewEntry()
	gui.guessEntry.SetPlaceHolder("🎤 Enter your guess here...")
	gui.guessEntry.OnSubmitted = func(text string) {
		gui.makeModernGuess()
	}

	gui.guessButton = widget.NewButton("🎯 Make Guess", gui.makeModernGuess)
	gui.newGameButton = widget.NewButton("🎮 New Game", gui.showModernCategorySelection)

	// GUI style switcher
	gui.switchGUIButton = widget.NewButton("🎤 Switch to Basic GUI", func() {
		manager := GetGUIManager()
		// Update game state in manager before switching
		if gui.game != nil {
			manager.UpdateGameState(gui.game, gui.selectedFile)
		}
		manager.SwitchToBasicGUI()
	})

	buttonContainer := container.NewGridWithColumns(3, gui.guessButton, gui.newGameButton, gui.switchGUIButton)

	gui.inputCard = container.NewVBox(
		widget.NewCard("💭 Your Guess", "", container.NewVBox(
			gui.guessEntry,
			widget.NewSeparator(),
			buttonContainer,
		)),
	)
}

func (gui *ModernGameGUI) createAdvancedWidgets() {
	// Create tabs for different views - start with empty tabs
	gui.tabs = container.NewAppTabs()

	// Add placeholder tabs that will be updated later
	gui.tabs.Append(container.NewTabItem("🎮 Game", widget.NewLabel("Loading...")))
	gui.tabs.Append(container.NewTabItem("📊 Stats", widget.NewLabel("Loading...")))
	gui.tabs.Append(container.NewTabItem("⚙️ Settings", widget.NewLabel("Loading...")))
}

func (gui *ModernGameGUI) setupThemeSelector() {
	gui.themeSelect = widget.NewSelect(gui.themeManager.GetThemeNames(), func(selected string) {
		gui.themeManager.SetTheme(selected)
	})
	gui.themeSelect.SetSelected("Hip-Hop")
}

func (gui *ModernGameGUI) showModernCategorySelection() {
	// Welcome section
	welcomeLabel := widget.NewLabel("Welcome to the ultimate hip-hop artist guessing experience!")
	welcomeLabel.Alignment = fyne.TextAlignCenter
	welcomeLabel.Wrapping = fyne.TextWrapWord

	instructionLabel := widget.NewLabel("Choose your category and start the musical journey:")
	instructionLabel.Alignment = fyne.TextAlignCenter

	// Category selection with enhanced display
	categoryOptions := make([]string, 0, len(CategoryNames))
	for i := 1; i <= 7; i++ {
		key := fmt.Sprintf("%d", i)
		name := CategoryNames[key]
		file := Categories[key]
		count := GetCategoryStats(file)
		option := fmt.Sprintf("%s: %s (%d artists)", key, name, count)
		categoryOptions = append(categoryOptions, option)
	}

	gui.categorySelect = widget.NewSelect(categoryOptions, gui.onModernCategorySelected)
	gui.categorySelect.PlaceHolder = "🎵 Choose your musical category..."

	// Create welcome card
	welcomeCard := widget.NewCard("🎵 Welcome", "", container.NewVBox(
		welcomeLabel,
		widget.NewSeparator(),
		instructionLabel,
		gui.categorySelect,
	))

	// Theme selection card
	themeCard := widget.NewCard("🎨 Themes", "", container.NewVBox(
		widget.NewLabel("Choose your visual style:"),
		gui.themeSelect,
	))

	// GUI style switcher card
	guiCard := widget.NewCard("🔄 GUI Style", "", container.NewVBox(
		widget.NewLabel("Switch interface style:"),
		gui.switchGUIButton,
	))

	// Side panel with additional options
	gui.sidePanel = container.NewVBox(
		themeCard,
		guiCard,
		widget.NewCard("ℹ️ Info", "", container.NewVBox(
			widget.NewLabel("Hip-Hop Hangman"),
			widget.NewLabel("Modern Edition"),
			widget.NewLabel("v2.0"),
		)),
	)

	// Main layout using border layout
	content := container.NewBorder(
		gui.toolbar,                      // Top
		nil,                              // Bottom
		gui.sidePanel,                    // Left
		nil,                              // Right
		container.NewCenter(welcomeCard), // Center
	)

	gui.window.SetContent(content)
}

func (gui *ModernGameGUI) onModernCategorySelected(selected string) {
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
		gui.startModernGame()
	}
}

func (gui *ModernGameGUI) startModernGame() {
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

	// Setup modern game UI
	gui.setupModernGameUI()
	gui.updateModernGameDisplay()
}

func (gui *ModernGameGUI) setupModernGameUI() {
	// Enable input
	gui.guessEntry.Enable()

	// Create main game layout using split container
	leftPanel := container.NewVBox(
		gui.hangmanCard,
		gui.statsCard,
	)

	rightPanel := container.NewVBox(
		gui.gameCard,
		gui.inputCard,
	)

	// Use split container for resizable panes
	splitContainer := container.NewHSplit(leftPanel, rightPanel)
	splitContainer.SetOffset(0.4) // 40% left, 60% right

	// Create tabs content
	gameTab := splitContainer
	statsTab := gui.createStatsTab()
	settingsTab := gui.createSettingsTab()

	// Update existing tabs with actual content
	gui.tabs.Items[0].Content = gameTab
	gui.tabs.Items[1].Content = statsTab
	gui.tabs.Items[2].Content = settingsTab

	// Refresh the tabs to show updated content
	gui.tabs.Refresh()

	// Main layout
	content := container.NewBorder(
		container.NewVBox(gui.titleCard, gui.toolbar), // Top
		nil,      // Bottom
		nil,      // Left
		nil,      // Right
		gui.tabs, // Center
	)

	gui.window.SetContent(content)
}

func (gui *ModernGameGUI) createStatsTab() *fyne.Container {
	// Create detailed statistics view
	gamesPlayedLabel := widget.NewLabel("Games Played: 1")
	winRateLabel := widget.NewLabel("Win Rate: 0%")
	avgScoreLabel := widget.NewLabel("Average Score: 0")
	bestScoreLabel := widget.NewLabel("Best Score: 0")

	statsCard := widget.NewCard("📈 Statistics", "", container.NewVBox(
		gamesPlayedLabel,
		winRateLabel,
		avgScoreLabel,
		bestScoreLabel,
	))

	// Create achievements section
	achievementsCard := widget.NewCard("🏆 Achievements", "", container.NewVBox(
		widget.NewLabel("🎯 First Win - Not yet achieved"),
		widget.NewLabel("🔥 Win Streak - Not yet achieved"),
		widget.NewLabel("💯 Perfect Game - Not yet achieved"),
	))

	return container.NewVBox(statsCard, achievementsCard)
}

func (gui *ModernGameGUI) createSettingsTab() *fyne.Container {
	// Create settings controls
	difficultySelect := widget.NewSelect([]string{"Easy", "Medium", "Hard"}, nil)
	difficultySelect.SetSelected("Medium")

	soundCheck := widget.NewCheck("Enable Sound Effects", nil)
	animationsCheck := widget.NewCheck("Enable Animations", nil)
	animationsCheck.SetChecked(true)

	settingsCard := widget.NewCard("⚙️ Game Settings", "", container.NewVBox(
		widget.NewLabel("Difficulty:"),
		difficultySelect,
		soundCheck,
		animationsCheck,
	))

	themeCard := widget.NewCard("🎨 Appearance", "", container.NewVBox(
		widget.NewLabel("Theme:"),
		gui.themeSelect,
	))

	return container.NewVBox(settingsCard, themeCard)
}

func (gui *ModernGameGUI) updateModernGameDisplay() {
	if gui.game == nil {
		return
	}

	// Update word display
	gui.wordLabel.SetText(gui.game.GetWordDisplay())

	// Update hangman art
	gui.hangmanLabel.SetText(GetEmojiHangmanArt(gui.game.Attempts))

	// Update stats
	gui.attemptsLabel.SetText(fmt.Sprintf("❤️ Lives: %d", gui.game.Attempts))
	gui.scoreLabel.SetText(fmt.Sprintf("🏆 Score: %d", gui.game.Score))

	// Update progress bar (based on remaining attempts)
	progress := float64(gui.game.Attempts) / 6.0
	gui.progressBar.SetValue(progress)

	// Update guessed letters
	if len(gui.game.GuessedLetters) > 0 {
		gui.guessedLabel.SetText(fmt.Sprintf("🔤 Guessed: %s", gui.game.GetGuessedLettersDisplay()))
	} else {
		gui.guessedLabel.SetText("🔤 No guesses yet")
	}

	// Check game over
	if gui.game.IsGameOver {
		gui.handleModernGameOver()
	}
}

func (gui *ModernGameGUI) makeModernGuess() {
	if gui.game == nil || gui.game.IsGameOver {
		return
	}

	guess := strings.TrimSpace(gui.guessEntry.Text)
	if guess == "" {
		gui.showModernDialog("Invalid Input", "Please enter a letter or word to guess.")
		return
	}

	// Check for duplicate single letter guesses
	if len(guess) == 1 {
		for _, g := range gui.game.GuessedLetters {
			if strings.ToUpper(g) == strings.ToUpper(guess) {
				gui.showModernDialog("Already Guessed", "You already guessed that letter!")
				gui.guessEntry.SetText("")
				return
			}
		}
	}

	// Make the guess
	correctGuess := gui.game.MakeGuess(guess)
	gui.guessEntry.SetText("")

	// Show feedback with modern styling
	if len(guess) == 1 {
		if !correctGuess {
			gui.showModernNotification("❌ Wrong letter!", "error")
		} else {
			gui.showModernNotification("✅ Good guess!", "success")
		}
	} else {
		if !correctGuess {
			gui.showModernNotification("❌ Wrong word!", "error")
		}
	}

	gui.updateModernGameDisplay()
}

func (gui *ModernGameGUI) showModernDialog(title, message string) {
	content := container.NewVBox(
		widget.NewLabel(message),
		widget.NewSeparator(),
	)

	dialog.ShowCustom(title, "OK", content, gui.window)
}

func (gui *ModernGameGUI) showModernNotification(message, notificationType string) {
	// Create a temporary notification overlay
	notification := widget.NewLabel(message)
	notification.Alignment = fyne.TextAlignCenter

	// This would ideally be a toast/snackbar, but we'll use a simple approach
	go func() {
		time.Sleep(2 * time.Second)
		// In a real implementation, we'd remove the notification here
	}()
}

func (gui *ModernGameGUI) handleModernGameOver() {
	gui.guessEntry.Disable()

	var title, message string
	if gui.game.IsWon {
		title = "🎉 Victory!"
		message = fmt.Sprintf("Congratulations! You guessed '%s'\nScore: %d points", gui.game.Word, gui.game.Score)
	} else {
		title = "💀 Game Over"
		message = fmt.Sprintf("The word was: %s\nBetter luck next time!", gui.game.Word)
	}

	dialog.ShowConfirm(title, message+"\n\nPlay again?", func(playAgain bool) {
		if playAgain {
			gui.showModernCategorySelection()
		} else {
			gui.app.Quit()
		}
	}, gui.window)
}

func (gui *ModernGameGUI) showSettingsDialog() {
	// Create settings dialog
	content := gui.createSettingsTab()
	dialog.ShowCustom("⚙️ Settings", "Close", content, gui.window)
}

func (gui *ModernGameGUI) showAboutDialog() {
	about := container.NewVBox(
		widget.NewLabel("Hip-Hop Hangman - Modern Edition"),
		widget.NewLabel("Version 2.0"),
		widget.NewSeparator(),
		widget.NewLabel("A modern take on the classic word game"),
		widget.NewLabel("featuring hip-hop artists and advanced UI."),
		widget.NewSeparator(),
		widget.NewLabel("Built with Fyne v2.6.3"),
	)

	dialog.ShowCustom("ℹ️ About", "Close", about, gui.window)
}
