package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// DemoComparisonGUI shows different FYNE design approaches side by side
type DemoComparisonGUI struct {
	app          fyne.App
	window       fyne.Window
	themeManager *ThemeManager

	// Demo games for comparison
	basicGame  *GameState
	modernGame *GameState
	customGame *GameState

	// UI Elements for different styles
	basicContainer  *fyne.Container
	modernContainer *fyne.Container
	customContainer *fyne.Container

	// Custom widgets
	customHangman    *CustomHangmanWidget
	animatedProgress *AnimatedProgressBar
	wordReveal       *WordRevealWidget
	scoreWidget      *ScoreWidget
}

func createDemoComparison() {
	demo := &DemoComparisonGUI{}
	demo.app = app.New()
	demo.themeManager = NewThemeManager(demo.app)
	demo.themeManager.SetTheme("Hip-Hop")

	demo.window = demo.app.NewWindow("🎨 FYNE Design Options Comparison 🎨")
	demo.window.Resize(fyne.NewSize(1400, 900))
	demo.window.SetMaster()

	demo.setupDemoUI()
	demo.window.ShowAndRun()
}

func (demo *DemoComparisonGUI) setupDemoUI() {
	// Initialize demo games with the same word for comparison
	word := "KENDRICK LAMAR"
	demo.basicGame = NewGame(word)
	demo.modernGame = NewGame(word)
	demo.customGame = NewGame(word)

	// Create different style demonstrations
	demo.createBasicStyleDemo()
	demo.createModernStyleDemo()
	demo.createCustomStyleDemo()

	// Create main layout with tabs for different comparisons
	tabs := container.NewAppTabs(
		container.NewTabItem("🔄 Side-by-Side Comparison", demo.createSideBySideComparison()),
		container.NewTabItem("🎨 Theme Showcase", demo.createThemeShowcase()),
		container.NewTabItem("🧩 Widget Gallery", demo.createWidgetGallery()),
		container.NewTabItem("📱 Layout Examples", demo.createLayoutExamples()),
		container.NewTabItem("🎯 Custom Components", demo.createCustomComponentsDemo()),
	)

	// Add header with information
	header := widget.NewCard("🎨 FYNE GUI Design Options", "", container.NewVBox(
		widget.NewLabel("This demo showcases the various design approaches available in FYNE:"),
		widget.NewLabel("• Basic widgets and standard layouts"),
		widget.NewLabel("• Modern Material Design-inspired interfaces"),
		widget.NewLabel("• Custom themes and color schemes"),
		widget.NewLabel("• Advanced widgets and animations"),
		widget.NewLabel("• Custom canvas-based components"),
	))

	content := container.NewBorder(header, nil, nil, nil, tabs)
	demo.window.SetContent(content)
}

func (demo *DemoComparisonGUI) createBasicStyleDemo() {
	// Basic FYNE approach (similar to current gui.go)
	titleLabel := widget.NewLabel("🎤 Basic Style")
	titleLabel.Alignment = fyne.TextAlignCenter
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	wordLabel := widget.NewLabel(demo.basicGame.GetWordDisplay())
	wordLabel.Alignment = fyne.TextAlignCenter
	wordLabel.TextStyle = fyne.TextStyle{Monospace: true}

	hangmanLabel := widget.NewLabel(GetEmojiHangmanArt(demo.basicGame.Attempts))
	hangmanLabel.Alignment = fyne.TextAlignCenter
	hangmanLabel.TextStyle = fyne.TextStyle{Monospace: true}

	attemptsLabel := widget.NewLabel(fmt.Sprintf("Lives: %d", demo.basicGame.Attempts))
	scoreLabel := widget.NewLabel(fmt.Sprintf("Score: %d", demo.basicGame.Score))

	guessEntry := widget.NewEntry()
	guessEntry.SetPlaceHolder("Enter guess...")
	guessButton := widget.NewButton("Guess", nil)

	demo.basicContainer = container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		container.NewGridWithColumns(2,
			container.NewVBox(hangmanLabel),
			container.NewVBox(wordLabel, attemptsLabel, scoreLabel),
		),
		widget.NewSeparator(),
		guessEntry,
		guessButton,
	)
}

func (demo *DemoComparisonGUI) createModernStyleDemo() {
	// Modern Material Design approach
	titleLabel := widget.NewLabel("🎨 Modern Style")
	titleLabel.Alignment = fyne.TextAlignCenter
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Use cards for modern look
	gameCard := widget.NewCard("🎯 Word", "", container.NewCenter(
		widget.NewLabel(demo.modernGame.GetWordDisplay()),
	))

	statsCard := widget.NewCard("📊 Stats", "", container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Score: %d", demo.modernGame.Score)),
		widget.NewLabel(fmt.Sprintf("Lives: %d", demo.modernGame.Attempts)),
		widget.NewProgressBar(), // Modern progress indicator
	))

	hangmanCard := widget.NewCard("🎭 Hangman", "", container.NewCenter(
		widget.NewLabel(GetEmojiHangmanArt(demo.modernGame.Attempts)),
	))

	inputCard := widget.NewCard("💭 Input", "", container.NewVBox(
		widget.NewEntry(),
		widget.NewButton("🎯 Make Guess", nil),
	))

	// Use split containers for modern layout
	topSplit := container.NewHSplit(hangmanCard, gameCard)
	bottomSplit := container.NewHSplit(statsCard, inputCard)

	demo.modernContainer = container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		topSplit,
		bottomSplit,
	)
}

func (demo *DemoComparisonGUI) createCustomStyleDemo() {
	// Custom widgets and advanced styling
	titleLabel := widget.NewLabel("🚀 Custom Style")
	titleLabel.Alignment = fyne.TextAlignCenter
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Use custom widgets
	demo.customHangman = NewCustomHangmanWidget()
	demo.customHangman.SetAttempts(demo.customGame.Attempts)

	demo.animatedProgress = NewAnimatedProgressBar()
	demo.animatedProgress.SetValue(float64(demo.customGame.Attempts) / 6.0)

	demo.wordReveal = NewWordRevealWidget(demo.customGame.Word)
	// Reveal some letters for demo
	demo.wordReveal.RevealLetter('K')
	demo.wordReveal.RevealLetter('E')

	demo.scoreWidget = NewScoreWidget()
	demo.scoreWidget.SetScore(demo.customGame.Score)

	customButton := NewCustomButton("🎯 Custom Guess", nil)

	demo.customContainer = container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		container.NewGridWithColumns(2,
			demo.customHangman,
			container.NewVBox(
				demo.wordReveal,
				demo.scoreWidget,
				demo.animatedProgress,
			),
		),
		widget.NewSeparator(),
		widget.NewEntry(),
		customButton,
	)
}

func (demo *DemoComparisonGUI) createSideBySideComparison() *fyne.Container {
	return container.NewGridWithColumns(3,
		widget.NewCard("Basic FYNE", "Standard widgets and layouts", demo.basicContainer),
		widget.NewCard("Modern Design", "Material Design inspired", demo.modernContainer),
		widget.NewCard("Custom Components", "Advanced custom widgets", demo.customContainer),
	)
}

func (demo *DemoComparisonGUI) createThemeShowcase() *fyne.Container {
	themes := []string{"Default Light", "Default Dark", "Hip-Hop", "Neon", "Retro", "Minimal"}

	themeCards := make([]*widget.Card, len(themes))
	for i, themeName := range themes {
		// Create a mini preview for each theme
		preview := container.NewVBox(
			widget.NewLabel("Sample Text"),
			widget.NewButton("Sample Button", nil),
			widget.NewEntry(),
		)

		themeCards[i] = widget.NewCard(themeName, "", container.NewVBox(
			preview,
			widget.NewButton("Apply Theme", func() {
				demo.themeManager.SetTheme(themeName)
			}),
		))
	}

	return container.NewGridWithColumns(3,
		themeCards[0], themeCards[1], themeCards[2],
		themeCards[3], themeCards[4], themeCards[5],
	)
}

func (demo *DemoComparisonGUI) createWidgetGallery() *fyne.Container {
	// Standard widgets
	standardWidgets := widget.NewCard("Standard Widgets", "", container.NewVBox(
		widget.NewLabel("Labels with different styles"),
		widget.NewButton("Buttons", nil),
		widget.NewEntry(),
		widget.NewCheck("Checkbox", nil),
		widget.NewSlider(0, 100),
		widget.NewProgressBar(),
		widget.NewSelect([]string{"Option 1", "Option 2"}, nil),
	))

	// Advanced widgets
	advancedWidgets := widget.NewCard("Advanced Widgets", "", container.NewVBox(
		widget.NewLabel("Rich Text and Lists"),
		widget.NewList(
			func() int { return 3 },
			func() fyne.CanvasObject { return widget.NewLabel("Item") },
			func(id widget.ListItemID, obj fyne.CanvasObject) {
				obj.(*widget.Label).SetText(fmt.Sprintf("List Item %d", id+1))
			},
		),
		widget.NewAccordion(
			widget.NewAccordionItem("Section 1", widget.NewLabel("Content 1")),
			widget.NewAccordionItem("Section 2", widget.NewLabel("Content 2")),
		),
	))

	// Custom widgets
	customWidgets := widget.NewCard("Custom Widgets", "", container.NewVBox(
		widget.NewLabel("Custom Canvas-based Components"),
		NewCustomHangmanWidget(),
		NewAnimatedProgressBar(),
		NewScoreWidget(),
	))

	return container.NewGridWithColumns(3, standardWidgets, advancedWidgets, customWidgets)
}

func (demo *DemoComparisonGUI) createLayoutExamples() *fyne.Container {
	// Border layout example
	borderExample := widget.NewCard("Border Layout", "", container.NewBorder(
		widget.NewLabel("North"),
		widget.NewLabel("South"),
		widget.NewLabel("West"),
		widget.NewLabel("East"),
		widget.NewLabel("Center"),
	))

	// Grid layout example
	gridExample := widget.NewCard("Grid Layout", "", container.NewGridWithColumns(2,
		widget.NewButton("1", nil),
		widget.NewButton("2", nil),
		widget.NewButton("3", nil),
		widget.NewButton("4", nil),
	))

	// VBox/HBox example
	boxExample := widget.NewCard("Box Layouts", "", container.NewVBox(
		widget.NewLabel("VBox Layout"),
		container.NewHBox(
			widget.NewButton("HBox 1", nil),
			widget.NewButton("HBox 2", nil),
			widget.NewButton("HBox 3", nil),
		),
	))

	// Split container example
	splitExample := widget.NewCard("Split Container", "",
		container.NewHSplit(
			widget.NewLabel("Left Panel"),
			widget.NewLabel("Right Panel"),
		),
	)

	// Scroll container example
	scrollContent := container.NewVBox()
	for i := 0; i < 20; i++ {
		scrollContent.Add(widget.NewLabel(fmt.Sprintf("Scrollable Item %d", i+1)))
	}
	scrollExample := widget.NewCard("Scroll Container", "",
		container.NewScroll(scrollContent),
	)

	// Tab container example
	tabExample := widget.NewCard("Tab Container", "",
		container.NewAppTabs(
			container.NewTabItem("Tab 1", widget.NewLabel("Content 1")),
			container.NewTabItem("Tab 2", widget.NewLabel("Content 2")),
			container.NewTabItem("Tab 3", widget.NewLabel("Content 3")),
		),
	)

	return container.NewGridWithColumns(3,
		borderExample, gridExample, boxExample,
		splitExample, scrollExample, tabExample,
	)
}

func (demo *DemoComparisonGUI) createCustomComponentsDemo() *fyne.Container {
	// Demonstrate custom canvas-based components
	hangmanDemo := widget.NewCard("Custom Hangman Widget", "Canvas-based drawing",
		NewCustomHangmanWidget(),
	)

	progressDemo := widget.NewCard("Animated Progress", "Smooth animations",
		NewAnimatedProgressBar(),
	)

	wordDemo := widget.NewCard("Word Reveal Widget", "Letter-by-letter reveal",
		NewWordRevealWidget("SAMPLE WORD"),
	)

	scoreDemo := widget.NewCard("Animated Score", "Number counting animation",
		NewScoreWidget(),
	)

	buttonDemo := widget.NewCard("Custom Button", "Hover and press effects",
		NewCustomButton("Custom Button", func() {
			dialog.ShowInformation("Custom Button", "Custom button clicked!", demo.window)
		}),
	)

	// Interactive demo controls
	controlsDemo := widget.NewCard("Interactive Controls", "", container.NewVBox(
		widget.NewButton("Test Hangman Animation", func() {
			demo.customHangman.SetAttempts(3) // Show partial hangman
		}),
		widget.NewButton("Test Progress Animation", func() {
			demo.animatedProgress.SetValue(0.7) // Animate to 70%
		}),
		widget.NewButton("Test Score Animation", func() {
			demo.scoreWidget.SetScore(1500) // Animate score change
		}),
		widget.NewButton("Reveal Letters", func() {
			demo.wordReveal.RevealLetter('S')
			demo.wordReveal.RevealLetter('A')
		}),
	))

	return container.NewGridWithColumns(3,
		hangmanDemo, progressDemo, wordDemo,
		scoreDemo, buttonDemo, controlsDemo,
	)
}

// Helper function to demonstrate different FYNE approaches
func ShowFyneDesignOptions() {
	options := []string{
		"1. Basic GUI (Current Implementation)",
		"2. Modern Material Design GUI",
		"3. Theme Comparison Demo",
		"4. Custom Widgets Demo",
		"5. Complete Design Comparison",
	}

	fmt.Println("🎨 FYNE GUI Design Options Available:")
	fmt.Println("=====================================")
	for _, option := range options {
		fmt.Println(option)
	}
	fmt.Println("\nTo see these in action, you can run:")
	fmt.Println("• createGUI() - Basic implementation")
	fmt.Println("• createModernGUI() - Modern design")
	fmt.Println("• createDemoComparison() - Full comparison")
}

// Function to create a simple theme switcher dialog
func createThemeSwitcher(app fyne.App, window fyne.Window) {
	themeManager := NewThemeManager(app)

	themeSelect := widget.NewSelect(themeManager.GetThemeNames(), func(selected string) {
		themeManager.SetTheme(selected)
	})
	themeSelect.SetSelected("Hip-Hop")

	content := container.NewVBox(
		widget.NewLabel("Choose a theme to see the visual changes:"),
		themeSelect,
		widget.NewSeparator(),
		widget.NewLabel("Themes demonstrate different design approaches:"),
		widget.NewLabel("• Hip-Hop: Custom brand colors"),
		widget.NewLabel("• Neon: Vibrant cyberpunk style"),
		widget.NewLabel("• Retro: 80s/90s aesthetic"),
		widget.NewLabel("• Minimal: Clean, simple design"),
	)

	dialog.ShowCustom("🎨 Theme Switcher", "Close", content, window)
}
