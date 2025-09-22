package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// HipHopTheme implements a custom theme for the Hip-Hop Hangman game
type HipHopTheme struct{}

// Color returns custom colors for different UI elements
func (h HipHopTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		if variant == theme.VariantLight {
			return color.RGBA{R: 245, G: 245, B: 250, A: 255} // Light purple-gray
		}
		return color.RGBA{R: 25, G: 25, B: 35, A: 255} // Dark navy
	case theme.ColorNameButton:
		return color.RGBA{R: 138, G: 43, B: 226, A: 255} // Purple
	case theme.ColorNamePrimary:
		return color.RGBA{R: 255, G: 215, B: 0, A: 255} // Gold
	case theme.ColorNameFocus:
		return color.RGBA{R: 255, G: 20, B: 147, A: 255} // Deep pink
	case theme.ColorNameHover:
		return color.RGBA{R: 186, G: 85, B: 211, A: 255} // Medium orchid
	case theme.ColorNameForeground:
		if variant == theme.VariantLight {
			return color.RGBA{R: 25, G: 25, B: 25, A: 255} // Dark text
		}
		return color.RGBA{R: 255, G: 255, B: 255, A: 255} // White text
	case theme.ColorNameSuccess:
		return color.RGBA{R: 50, G: 205, B: 50, A: 255} // Lime green
	case theme.ColorNameError:
		return color.RGBA{R: 220, G: 20, B: 60, A: 255} // Crimson
	case theme.ColorNameWarning:
		return color.RGBA{R: 255, G: 165, B: 0, A: 255} // Orange
	}

	// Fall back to default theme for other colors
	return theme.DefaultTheme().Color(name, variant)
}

// Font returns custom fonts for different text elements
func (h HipHopTheme) Font(style fyne.TextStyle) fyne.Resource {
	// Use default fonts but could be customized with custom font resources
	return theme.DefaultTheme().Font(style)
}

// Icon returns custom icons
func (h HipHopTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	// Use default icons but could be customized
	return theme.DefaultTheme().Icon(name)
}

// Size returns custom sizes for UI elements
func (h HipHopTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 16 // Larger default text
	case theme.SizeNameCaptionText:
		return 14
	case theme.SizeNameHeadingText:
		return 24
	case theme.SizeNameSubHeadingText:
		return 20
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameInlineIcon:
		return 24
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameSeparatorThickness:
		return 2
	}

	return theme.DefaultTheme().Size(name)
}

// NeonTheme - A vibrant neon-style theme
type NeonTheme struct{}

func (n NeonTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 10, G: 10, B: 20, A: 255} // Very dark blue
	case theme.ColorNameButton:
		return color.RGBA{R: 0, G: 255, B: 255, A: 255} // Cyan
	case theme.ColorNamePrimary:
		return color.RGBA{R: 255, G: 0, B: 255, A: 255} // Magenta
	case theme.ColorNameFocus:
		return color.RGBA{R: 0, G: 255, B: 0, A: 255} // Lime
	case theme.ColorNameHover:
		return color.RGBA{R: 255, G: 255, B: 0, A: 255} // Yellow
	case theme.ColorNameForeground:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255} // White
	case theme.ColorNameSuccess:
		return color.RGBA{R: 0, G: 255, B: 127, A: 255} // Spring green
	case theme.ColorNameError:
		return color.RGBA{R: 255, G: 20, B: 147, A: 255} // Deep pink
	case theme.ColorNameWarning:
		return color.RGBA{R: 255, G: 165, B: 0, A: 255} // Orange
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (n NeonTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (n NeonTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (n NeonTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 18 // Even larger text for neon theme
	case theme.SizeNameHeadingText:
		return 28
	case theme.SizeNameSubHeadingText:
		return 22
	case theme.SizeNamePadding:
		return 12 // More padding for neon glow effect
	}

	return theme.DefaultTheme().Size(name)
}

// RetroTheme - A retro 80s/90s style theme
type RetroTheme struct{}

func (r RetroTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 40, G: 44, B: 52, A: 255} // Dark gray
	case theme.ColorNameButton:
		return color.RGBA{R: 255, G: 105, B: 180, A: 255} // Hot pink
	case theme.ColorNamePrimary:
		return color.RGBA{R: 255, G: 215, B: 0, A: 255} // Gold
	case theme.ColorNameFocus:
		return color.RGBA{R: 0, G: 191, B: 255, A: 255} // Deep sky blue
	case theme.ColorNameHover:
		return color.RGBA{R: 255, G: 20, B: 147, A: 255} // Deep pink
	case theme.ColorNameForeground:
		return color.RGBA{R: 248, G: 248, B: 242, A: 255} // Off-white
	case theme.ColorNameSuccess:
		return color.RGBA{R: 50, G: 205, B: 50, A: 255} // Lime green
	case theme.ColorNameError:
		return color.RGBA{R: 255, G: 69, B: 0, A: 255} // Red orange
	case theme.ColorNameWarning:
		return color.RGBA{R: 255, G: 140, B: 0, A: 255} // Dark orange
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (r RetroTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (r RetroTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (r RetroTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

// MinimalTheme - A clean, minimal design theme
type MinimalTheme struct{}

func (m MinimalTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		if variant == theme.VariantLight {
			return color.RGBA{R: 255, G: 255, B: 255, A: 255} // Pure white
		}
		return color.RGBA{R: 18, G: 18, B: 18, A: 255} // Almost black
	case theme.ColorNameButton:
		if variant == theme.VariantLight {
			return color.RGBA{R: 70, G: 70, B: 70, A: 255} // Dark gray
		}
		return color.RGBA{R: 200, G: 200, B: 200, A: 255} // Light gray
	case theme.ColorNamePrimary:
		return color.RGBA{R: 100, G: 100, B: 100, A: 255} // Medium gray
	case theme.ColorNameFocus:
		return color.RGBA{R: 0, G: 122, B: 255, A: 255} // System blue
	case theme.ColorNameHover:
		return color.RGBA{R: 150, G: 150, B: 150, A: 255} // Light gray
	case theme.ColorNameForeground:
		if variant == theme.VariantLight {
			return color.RGBA{R: 0, G: 0, B: 0, A: 255} // Black
		}
		return color.RGBA{R: 255, G: 255, B: 255, A: 255} // White
	case theme.ColorNameSuccess:
		return color.RGBA{R: 52, G: 199, B: 89, A: 255} // System green
	case theme.ColorNameError:
		return color.RGBA{R: 255, G: 59, B: 48, A: 255} // System red
	case theme.ColorNameWarning:
		return color.RGBA{R: 255, G: 149, B: 0, A: 255} // System orange
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m MinimalTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m MinimalTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m MinimalTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 4 // Minimal padding
	case theme.SizeNameSeparatorThickness:
		return 1 // Thin separators
	}

	return theme.DefaultTheme().Size(name)
}

// ThemeManager handles switching between different themes
type ThemeManager struct {
	app    fyne.App
	themes map[string]fyne.Theme
}

func NewThemeManager(app fyne.App) *ThemeManager {
	tm := &ThemeManager{
		app:    app,
		themes: make(map[string]fyne.Theme),
	}

	// Register available themes
	tm.themes["Default Light"] = theme.LightTheme()
	tm.themes["Default Dark"] = theme.DarkTheme()
	tm.themes["Hip-Hop"] = &HipHopTheme{}
	tm.themes["Neon"] = &NeonTheme{}
	tm.themes["Retro"] = &RetroTheme{}
	tm.themes["Minimal"] = &MinimalTheme{}

	return tm
}

func (tm *ThemeManager) GetThemeNames() []string {
	names := make([]string, 0, len(tm.themes))
	for name := range tm.themes {
		names = append(names, name)
	}
	return names
}

func (tm *ThemeManager) SetTheme(name string) {
	if theme, exists := tm.themes[name]; exists {
		tm.app.Settings().SetTheme(theme)
	}
}

func (tm *ThemeManager) GetTheme(name string) fyne.Theme {
	return tm.themes[name]
}
