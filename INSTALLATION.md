# 🎤 Hip-Hop Hangman - Installation Guide

## Quick Start (Recommended)

### Prerequisites
- **Go 1.19+** installed on your system
- **Git** for cloning the repository

### Installation Steps

1. **Clone the repository:**
   ```bash
   git clone https://github.com/jamesBoder/hiphop_hangman.git
   cd hiphop_hangman
   ```

2. **Switch to Version 3 (Dynamic GUI Switching):**
   ```bash
   git checkout version3
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Build the application:**
   ```bash
   go build
   ```

5. **Run the game:**
   ```bash
   ./hiphop_hangman
   ```

## 🎮 How to Play

### Starting the Game
- **Default**: `./hiphop_hangman` (starts with Basic GUI)
- **Specific GUI**: `./hiphop_hangman basic` or `./hiphop_hangman modern`
- **Demo Mode**: `./hiphop_hangman demo`

### Dynamic GUI Switching (NEW!)
- Click **"🎨 Switch to Modern GUI"** in Basic mode
- Click **"🎤 Switch to Basic GUI"** in Modern mode
- Your game progress is preserved when switching!

### Theme Selection
- Choose from 6 themes: Hip-Hop, Neon, Retro, Minimal, Light, Dark
- Available in both Basic and Modern GUIs
- Change themes anytime without restarting

## 🛠️ System Requirements

### Desktop Platforms (Fully Supported)
- **OS**: Windows 10+, macOS 10.14+, or Linux (Ubuntu 18.04+)
- **RAM**: 512MB available memory
- **Go**: Version 1.19 or higher
- **Display**: GUI support required

### Mobile Platforms (Android/iOS)
**⚠️ Current Status: Limited Support**

**Android:**
- **Possible**: Yes, with Fyne's mobile support
- **Requires**: Additional setup with `fyne package -os android`
- **Limitations**: 
  - Touch interface not optimized (designed for desktop)
  - Small screen may affect readability
  - Requires Android SDK and development setup
  - No direct APK distribution currently available

**iOS:**
- **Possible**: Yes, with Fyne's mobile support  
- **Requires**: macOS with Xcode and `fyne package -os ios`
- **Limitations**: Same as Android plus Apple Developer account needed

### Mobile Installation (Advanced Users)
```bash
# For Android (requires Android SDK)
fyne package -os android -appID com.example.hiphophangman

# For iOS (requires Xcode on macOS)
fyne package -os ios -appID com.example.hiphophangman
```

**Recommendation**: Currently optimized for desktop use. Mobile versions would need UI redesign for optimal touch experience.

### Dependencies (Auto-installed)
- Fyne v2.6.3 (GUI framework with mobile support)
- Go standard library

## 📋 Alternative Installation Methods

### Method 1: Direct Download
1. Go to: https://github.com/jamesBoder/hiphop_hangman
2. Click **"Code"** → **"Download ZIP"**
3. Extract the ZIP file
4. Open terminal in the extracted folder
5. Run: `git checkout version3` (if git is available) or manually switch to version3 branch
6. Follow steps 3-5 from Quick Start

### Method 2: Go Install (if published)
```bash
go install github.com/jamesBoder/hiphop_hangman@version3
```

## 🚨 Troubleshooting

### Common Issues

**"Go not found" error:**
- Install Go from: https://golang.org/dl/
- Add Go to your system PATH

**"Permission denied" error (Linux/Mac):**
```bash
chmod +x hiphop_hangman
./hiphop_hangman
```

**GUI doesn't appear:**
- Ensure you have a desktop environment
- On Linux, install: `sudo apt-get install libgl1-mesa-dev xorg-dev`

**Build fails:**
```bash
go clean -modcache
go mod download
go build
```

## 🎯 Features Overview

### Version 3 Highlights
- **Dynamic GUI Switching**: Switch between Basic and Modern interfaces while playing
- **Game State Preservation**: Never lose progress when switching GUIs
- **6 Custom Themes**: Hip-Hop, Neon, Retro, Minimal, Light, Dark
- **Hip-Hop Categories**: East Coast, West Coast, South, Midwest, International, Groups
- **3000+ Artists**: Comprehensive database of hip-hop artists

### GUI Modes
- **Basic GUI**: Clean, simple interface with essential features
- **Modern GUI**: Advanced Material Design with cards, tabs, and animations
- **Demo Mode**: Side-by-side comparison of all features

## 📞 Support

### Getting Help
- **Issues**: https://github.com/jamesBoder/hiphop_hangman/issues
- **Discussions**: https://github.com/jamesBoder/hiphop_hangman/discussions

### Contributing
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

---

**Enjoy the game! 🎤🎮**
