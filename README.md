# Command Line Timer

A beautifully crafted command line countdown timer that combines minimalist design with powerful functionality. Perfect for productivity, cooking, workouts, or any timed activity.

![Timer Demo](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Unix-blue?style=for-the-badge)

## ✨ Key Features

### 🎯 **Perfectly Centered Display**
- Auto-detects terminal dimensions and centers the timer perfectly
- Unicode-aware positioning ensures consistent alignment across all terminals

### 🔢 **Large Digital Display**
- Clean, block-style digital font using Unicode characters
- Bold white text for maximum readability
- Optimized character spacing for visual clarity

### ⏰ **Consistent Time Format**
- Always displays time in `HH:MM:SS` format
- Handles any duration from seconds to hours seamlessly
- Easy-to-read countdown with smooth updates

### ✨ **True Blinking Alert**
- Genuine show/hide blinking effect when timer expires
- 5-second alert cycle with 200ms intervals
- Bright, attention-grabbing visual feedback

### 🎮 **Smart Keyboard Controls**
- **Ctrl+Q**: Graceful exit during countdown (prevents accidental termination)
- **Enter**: Stop blinking when timer expires
- **Ctrl+C**: Immediate emergency exit (always available)

### 🖥️ **Cross-Platform Compatibility**
- Native builds for Linux, macOS (Intel and ARM64)
- Works on any Unix-like system with Go support
- No external dependencies beyond standard libraries

## 🚀 Quick Start

### Installation Options

#### Method 1: From Source (Recommended)
```bash
git clone https://github.com/yourusername/commandline-timer.git
cd commandline-timer
make build
sudo make install  # Optional: installs to /usr/local/bin
```

#### Method 2: Manual Build
```bash
go build -o timer ./cmd/timer/main.go
chmod +x timer
cp timer /usr/local/bin/  # Optional: system-wide installation
```

#### Method 3: Using Build Script
```bash
# Build for current platform
./build.sh

# Build for all platforms
./build.sh --all
```

## 📖 Usage Guide

### Basic Usage
```bash
# Simple timer for 30 seconds
timer 30s

# 5 minute timer
timer 5m

# 1 hour and 30 minutes
timer 1h30m

# Complex duration: 2 hours, 45 minutes, 30 seconds
timer 2h45m30s
```

### Duration Format Support
- **Seconds**: `30s`, `45s`, `60s`
- **Minutes**: `5m`, `10m`, `30m`
- **Hours**: `1h`, `2h`, `24h`
- **Combined**: `1h30m`, `2h45m30s`, `1h5m30s`

### During Timer Operation
1. **Timer Display**: Large, centered countdown in `HH:MM:SS` format
2. **Progress**: Updates every 100ms for smooth visual feedback
3. **Instructions**: Help text displayed at bottom of screen
4. **Controls**: Responsive keyboard input for immediate action

### When Timer Expires
1. **Alert**: "00:00:00" displays with true blinking effect
2. **Duration**: 5-second blinking cycle
3. **Stop**: Press Enter to stop blinking and exit
4. **Message**: "Time's up!" confirmation in terminal

## 🛠️ Technical Details

### Architecture
- **Language**: Go 1.21+ for performance and cross-platform compilation
- **Dependencies**: Minimal external libraries
  - `github.com/mattn/go-runewidth`: Unicode text width calculation
  - `golang.org/x/term`: Terminal operations and raw mode handling
- **Design**: Single-file application with clear separation of concerns

### Key Components
- **Timer Core**: Precise duration tracking with millisecond accuracy
- **Display Engine**: Terminal-aware rendering with Unicode support
- **Input Handler**: Non-blocking keyboard input processing
- **Signal Management**: Graceful handling of system signals

### Terminal Integration
- **Raw Mode**: Direct terminal control for immediate input response
- **Cursor Management**: Automatic hiding/showing for clean display
- **Color Support**: ANSI color codes for visual enhancement
- **Screen Management**: Proper cleanup and restoration on exit

## 🏗️ Building and Development

### Build Commands
```bash
# Build for current platform
make build

# Build for specific platforms
make build-linux      # Linux AMD64
make build-macos      # macOS Intel
make build-macos-arm64 # macOS ARM64

# Build for all platforms
make build-all

# Clean build artifacts
make clean

# Download dependencies
make deps

# Run tests
make test
```

### Development Setup
```bash
# Clone repository
git clone https://github.com/yourusername/commandline-timer.git
cd commandline-timer

# Download dependencies
go mod download
go mod tidy

# Build and test
make build
make test

# Install locally
make build && ./timer 1m  # Test with 1-minute timer
```

## 🎮 Controls Reference

| Key | Action | When Available |
|-----|--------|----------------|
| `Ctrl+Q` | Graceful exit | During countdown |
| `Enter` | Stop blinking | When timer expired |
| `Ctrl+C` | Immediate exit | Always available |

## 📋 Requirements

- **Go**: Version 1.21 or higher
- **Terminal**: Any modern terminal emulator
- **OS**: Linux, macOS, or Unix-like system
- **Memory**: Minimal footprint (< 10MB RAM)
- **Disk**: ~5MB for compiled binary

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

### Development Workflow
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with Go for its simplicity and performance
- Inspired by the need for a clean, functional terminal timer
- Thanks to the Go community for excellent libraries and tools

---

**Made with ❤️ for command line enthusiasts**