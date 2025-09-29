package main

import (
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// GUIStyle represents the different GUI implementation styles
type GUIStyle int

const (
	BasicGUI GUIStyle = iota
	ModernGUI
	DemoGUI
)

// UnifiedGUIManager manages switching between different GUI styles while preserving game state
type UnifiedGUIManager struct {
	app          fyne.App
	window       fyne.Window
	currentStyle GUIStyle
	themeManager *ThemeManager

	// Shared game state
	gameState          *GameState
	selectedFile       string
	selectedDifficulty Difficulty

	// GUI instances
	basicGUI  *GameGUI
	modernGUI *ModernGameGUI
}

// NewUnifiedGUIManager creates a new unified GUI manager
func NewUnifiedGUIManager() *UnifiedGUIManager {
	manager := &UnifiedGUIManager{
		app:          app.New(),
		currentStyle: BasicGUI, // Start with basic GUI
	}

	manager.window = manager.app.NewWindow("🎤 Hip-Hop Hangman 🎤")
	manager.window.Resize(fyne.NewSize(900, 700))

	// Initialize theme manager
	manager.themeManager = NewThemeManager(manager.app)
	manager.themeManager.SetTheme("Hip-Hop") // Default theme

	return manager
}

// SwitchToBasicGUI switches to the basic GUI style
func (m *UnifiedGUIManager) SwitchToBasicGUI() {
	fmt.Println("🎤 Switching to Basic GUI...")
	m.currentStyle = BasicGUI

	// Create or reuse basic GUI instance
	if m.basicGUI == nil {
		m.basicGUI = &GameGUI{
			app:          m.app,
			window:       m.window,
			themeManager: m.themeManager, // Add theme support
		}
		m.basicGUI.setupUI()
	}

	// Transfer game state if exists
	if m.gameState != nil {
		m.basicGUI.game = m.gameState
		m.basicGUI.selectedFile = m.selectedFile
		m.basicGUI.selectedDifficulty = m.selectedDifficulty
		m.basicGUI.setupGameUI()
		m.basicGUI.updateGameDisplay()
	} else {
		m.basicGUI.showCategorySelection()
	}
}

// SwitchToModernGUI switches to the modern GUI style
func (m *UnifiedGUIManager) SwitchToModernGUI() {
	fmt.Println("🎨 Switching to Modern GUI...")
	m.currentStyle = ModernGUI

	// Create or reuse modern GUI instance
	if m.modernGUI == nil {
		m.modernGUI = &ModernGameGUI{
			app:          m.app,
			window:       m.window,
			themeManager: m.themeManager,
		}
		m.modernGUI.setupModernUI()
	}

	// Transfer game state if exists
	if m.gameState != nil {
		m.modernGUI.game = m.gameState
		m.modernGUI.selectedFile = m.selectedFile
		m.modernGUI.selectedDifficulty = m.selectedDifficulty
		m.modernGUI.setupModernGameUI()
		m.modernGUI.updateModernGameDisplay()
	} else {
		m.modernGUI.showModernCategorySelection()
	}
}

// SwitchToDemoGUI switches to the demo comparison GUI
func (m *UnifiedGUIManager) SwitchToDemoGUI() {
	fmt.Println("🚀 Switching to Demo GUI...")
	m.currentStyle = DemoGUI
	createDemoComparison()
}

// UpdateGameState updates the shared game state
func (m *UnifiedGUIManager) UpdateGameState(game *GameState, selectedFile string) {
	m.gameState = game
	m.selectedFile = selectedFile
	if game != nil {
		m.selectedDifficulty = game.Difficulty
	}
}

// GetCurrentStyle returns the current GUI style
func (m *UnifiedGUIManager) GetCurrentStyle() GUIStyle {
	return m.currentStyle
}

// GetThemeManager returns the theme manager
func (m *UnifiedGUIManager) GetThemeManager() *ThemeManager {
	return m.themeManager
}

// Run starts the unified GUI manager
func (m *UnifiedGUIManager) Run() {
	// Check for command line arguments for backward compatibility
	if len(os.Args) > 1 {
		mode := os.Args[1]
		switch mode {
		case "basic":
			m.SwitchToBasicGUI()
		case "modern":
			m.SwitchToModernGUI()
		case "demo":
			m.SwitchToDemoGUI()
		case "help":
			showGUIHelp()
			return
		default:
			fmt.Printf("Unknown mode: %s\n", mode)
			showGUIHelp()
			return
		}
	} else {
		// Start with basic GUI by default, but allow switching
		m.SwitchToBasicGUI()
	}

	m.window.ShowAndRun()
}

// Global manager instance
var globalGUIManager *UnifiedGUIManager

// GetGUIManager returns the global GUI manager instance
func GetGUIManager() *UnifiedGUIManager {
	if globalGUIManager == nil {
		globalGUIManager = NewUnifiedGUIManager()
	}
	return globalGUIManager
}

// Legacy function for backward compatibility - now uses unified manager
func selectGUIMode() {
	manager := GetGUIManager()
	manager.Run()
}

func showGUIHelp() {
	fmt.Println("🎨 FYNE GUI Design Options")
	fmt.Println("==========================")
	fmt.Println("Usage: ./hiphop_hangman [mode]")
	fmt.Println("")
	fmt.Println("Available modes:")
	fmt.Println("  basic   - Basic FYNE implementation (default)")
	fmt.Println("  modern  - Modern Material Design interface")
	fmt.Println("  demo    - Comprehensive design comparison")
	fmt.Println("  help    - Show this help message")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  ./hiphop_hangman basic")
	fmt.Println("  ./hiphop_hangman modern")
	fmt.Println("  ./hiphop_hangman demo")
	fmt.Println("")
	fmt.Println("NEW: You can now switch between GUI styles while playing!")
	fmt.Println("Look for the 'Switch GUI Style' option in the interface.")
}

// Alternative function to run specific demos programmatically
func runGUIDemo(demoType string) {
	switch demoType {
	case "themes":
		fmt.Println("🎨 Theme Demo - showing different visual styles...")
		createDemoComparison()
	case "widgets":
		fmt.Println("🧩 Widget Demo - showing custom components...")
		createDemoComparison()
	case "layouts":
		fmt.Println("📱 Layout Demo - showing different arrangements...")
		createDemoComparison()
	default:
		fmt.Printf("Unknown demo type: %s\n", demoType)
		ShowFyneDesignOptions()
	}
}

// Function to demonstrate theme switching programmatically
func demonstrateThemes() {
	themes := []string{"Hip-Hop", "Neon", "Retro", "Minimal", "Default Light", "Default Dark"}

	fmt.Println("🎨 Available Themes:")
	fmt.Println("===================")
	for i, theme := range themes {
		fmt.Printf("%d. %s\n", i+1, theme)
	}
	fmt.Println("")
	fmt.Print("Enter theme number to preview (1-6): ")

	var choice string
	fmt.Scanln(&choice)

	if num, err := strconv.Atoi(choice); err == nil && num >= 1 && num <= len(themes) {
		selectedTheme := themes[num-1]
		fmt.Printf("🎨 Starting demo with %s theme...\n", selectedTheme)
		createDemoComparison()
	} else {
		fmt.Printf("Invalid choice: %s\n", choice)
		demonstrateThemes()
	}
}

// Quick function to show what's possible with FYNE
func showFyneCapabilities() {
	fmt.Println("🚀 FYNE GUI Framework Capabilities")
	fmt.Println("===================================")
	fmt.Println("")
	fmt.Println("📊 Layout Options:")
	fmt.Println("  • VBox/HBox - Simple stacking")
	fmt.Println("  • Grid - Organized columns/rows")
	fmt.Println("  • Border - North/South/East/West/Center")
	fmt.Println("  • Split - Resizable panes")
	fmt.Println("  • Tabs - Multiple views")
	fmt.Println("  • Scroll - Overflow handling")
	fmt.Println("")
	fmt.Println("🎨 Theming Options:")
	fmt.Println("  • Built-in Light/Dark themes")
	fmt.Println("  • Custom color schemes")
	fmt.Println("  • Brand-specific styling")
	fmt.Println("  • Dynamic theme switching")
	fmt.Println("")
	fmt.Println("🧩 Widget Options:")
	fmt.Println("  • Standard: Labels, Buttons, Entries, etc.")
	fmt.Println("  • Advanced: Cards, Lists, Accordions, etc.")
	fmt.Println("  • Custom: Canvas-based drawing and animations")
	fmt.Println("")
	fmt.Println("✨ Advanced Features:")
	fmt.Println("  • Smooth animations (60 FPS)")
	fmt.Println("  • Mouse hover and interaction effects")
	fmt.Println("  • Custom canvas-based graphics")
	fmt.Println("  • Responsive layouts")
	fmt.Println("  • Cross-platform compatibility")
	fmt.Println("")
	fmt.Println("NEW: Dynamic GUI style switching during gameplay!")
	fmt.Println("Switch between Basic and Modern styles anytime! 🎯")
}
