# Command Line Timer

A beautifully crafted command line countdown timer that combines minimalist design with powerful functionality. Perfect for productivity, cooking, workouts, or any timed activity.

![Timer Demo](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Unix-blue?style=for-the-badge)

## ✨ Features

- 🎯 **Perfectly Centered Display** - Auto-detects terminal dimensions and centers perfectly
- 🔢 **Large Digital Display** - Clean block-style font with bold white text
- ⏰ **Consistent Time Format** - Always displays in `HH:MM:SS` format
- ✨ **True Blinking Alert** - Genuine show/hide blinking when timer expires
- 🎮 **Smart Controls** - Ctrl+Q to exit, Enter to stop blinking
- 🖥️ **Cross-Platform** - Works on Linux, macOS, and Unix-like systems

## 🚀 Installation

### From Source
```bash
git clone https://github.com/yourusername/commandline-timer.git
cd commandline-timer
make build
sudo make install  # Optional: installs to /usr/local/bin
```

### Manual Build
```bash
go build -o timer ./cmd/timer/main.go
chmod +x timer
cp timer /usr/local/bin/  # Optional: system-wide installation
```

### Build Script
```bash
./build.sh        # Current platform
./build.sh --all   # All platforms
```

## 📖 Usage

```bash
timer 30s          # 30 seconds
timer 5m           # 5 minutes
timer 1h30m        # 1 hour 30 minutes
timer 2h45m30s     # 2 hours, 45 minutes, 30 seconds
```

### Duration Format Support
- **Seconds**: `30s`, `45s`, `60s`
- **Minutes**: `5m`, `10m`, `30m`
- **Hours**: `1h`, `2h`, `24h`
- **Combined**: `1h30m`, `2h45m30s`, `1h5m30s`

### Controls
| Key | Action | When Available |
|-----|--------|----------------|
| `Ctrl+Q` | Graceful exit | During countdown |
| `Enter` | Stop blinking | When timer expired |
| `Ctrl+C` | Immediate exit | Always available |

## 🛠️ Technical Details

Built with Go 1.21+ using minimal dependencies:
- `github.com/mattn/go-runewidth` - Unicode text width calculation
- `golang.org/x/term` - Terminal operations and raw mode handling

### Key Features
- **Precise Timing**: Millisecond accuracy with smooth 100ms updates
- **Terminal Integration**: Raw mode for immediate input response
- **Unicode Support**: Consistent display across all terminals
- **Clean Architecture**: Single-file application with clear separation of concerns

## 🏗️ Development

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Clean build artifacts
make clean

# Run tests
make test
```

## 📋 Requirements

- Go 1.21+
- Linux, macOS, or Unix-like system

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Made with ❤️ for command line enthusiasts**