# Command Line Timer

A minimalist command line countdown timer with perfectly centered digital display and visual alerts.

## Features

- 🎯 **Perfectly Centered**: Auto-centered display using Unicode-aware positioning
- 🔢 **Large Digital Display**: Clean block-style numbers for maximum readability
- ⏰ **HH:MM:SS Format**: Consistent time display format for all durations
- ✨ **True Blinking Alert**: Genuine show/hide blinking when timer expires
- 🎮 **Smart Controls**: Ctrl+Q to exit (prevents accidental termination)
- 🖥️ **Cross-Platform**: Works on Linux, macOS, and Unix-like systems

## Installation

### From Source

```bash
git clone https://github.com/yourusername/commandline-timer.git
cd commandline-timer
make build
sudo make install  # Optional
```

### Manual Build

```bash
go build -o timer ./cmd/timer/main.go
chmod +x timer
cp timer /usr/local/bin/  # Optional
```

## Usage

```bash
timer 30s     # 30 seconds
timer 5m      # 5 minutes  
timer 1h30m   # 1 hour 30 minutes
timer 2h45m30s # 2 hours, 45 minutes, 30 seconds
```

## Controls

- **Ctrl+Q**: Exit during countdown
- **Enter**: Stop blinking when timer expires
- **Ctrl+C**: Immediate exit (always available)

## Technical Details

Built with Go 1.21+ using minimal dependencies:
- `github.com/mattn/go-runewidth` - Unicode text width calculation
- `golang.org/x/term` - Terminal operations

The timer detects terminal dimensions and centers the display perfectly. When time expires, it shows "00:00:00" with true blinking (5 seconds at 200ms intervals).

## Building

```bash
make build         # Current platform
make build-all     # All platforms
make clean         # Clean artifacts
```

## License

MIT License