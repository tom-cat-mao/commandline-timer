# Command Line Timer

A minimalist yet powerful command line countdown timer written in Go. Features a perfectly centered digital display with large, clean numbers and smooth visual feedback.

## Features

- 🎯 **Perfectly Centered Display**: Timer is automatically centered both horizontally and vertically
- 🔢 **Large Digital Display**: Clean, block-style digital numbers for maximum readability
- ⏰ **HH:MM:SS Format**: Always displays time in hours:minutes:seconds format
- ✨ **True Blinking Alert**: Genuine blinking effect when timer expires (00:00:00 flashes for 5 seconds)
- 🎨 **Clean Interface**: Bold white text on dark background for optimal contrast
- 🖥️ **Cross-Platform**: Works on Linux, macOS, and other Unix-like systems
- 🚀 **Lightweight**: Minimal dependencies, fast startup
- 🎮 **Smart Controls**: Prevents accidental termination during countdown

## Key Features Explained

### Centered Display
The timer automatically detects terminal dimensions and centers the display perfectly using Unicode-aware width calculation for accurate positioning.

### Digital Display
- Uses custom block-style font for large, readable numbers
- Each digit is 5 characters tall with proper spacing
- Maintains perfect alignment and centering

### Time Format
- Always displays in HH:MM:SS format regardless of duration
- Examples: 00:00:30 (30 seconds), 00:05:00 (5 minutes), 01:30:00 (1.5 hours)

### Blinking Alert
When time expires:
- Displays "00:00:00" with true blinking (show/hide alternation)
- Blinks for exactly 5 seconds at 200ms intervals
- Clean screen clear when finished

### Smart Controls
- **During Countdown**: Only Ctrl+Q can exit (prevents accidental termination)
- **During Final Blink**: Enter key can stop the blinking
- **Always**: Ctrl+C for immediate exit

## Installation

### From Source

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/commandline-timer.git
   cd commandline-timer
   ```

2. Build the binary:
   ```bash
   make build
   ```

3. Install system-wide (optional):
   ```bash
   sudo make install
   ```

### Manual Installation

1. Build the binary:
   ```bash
   go build -o timer ./main.go
   ```

2. Make it executable:
   ```bash
   chmod +x timer
   ```

3. Copy to your PATH:
   ```bash
   cp timer /usr/local/bin/
   ```

## Usage

### Basic Usage

```bash
timer 30s    # 30 seconds (displays as 00:00:30)
timer 5m     # 5 minutes (displays as 00:05:00)
timer 1h30m  # 1 hour 30 minutes (displays as 01:30:00)
timer 2h45m30s  # 2 hours, 45 minutes, 30 seconds (displays as 02:45:30)
```

### Examples

```bash
# Quick 10-second countdown
timer 10s

# 5-minute work timer
timer 5m

# 1-hour break timer
timer 1h

# Complex duration
timer 1h30m15s
```

### Controls

- **Ctrl+Q**: Exit during countdown (prevents accidental termination)
- **Enter**: Stop blinking during final alert
- **Ctrl+C**: Immediate exit (always available)

## Building

### Build for Current Platform

```bash
make build
```

### Build for All Platforms

```bash
make build-all
```

### Clean Build Artifacts

```bash
make clean
```

## Dependencies

- Go 1.21 or higher
- The following Go packages (automatically downloaded):
  - `github.com/mattn/go-runewidth` - For Unicode-aware text width calculation
  - `golang.org/x/term` - For terminal operations

## Technical Details

### Core Components

1. **Timer Logic**: Handles countdown calculations and time formatting
2. **Display Engine**: Manages terminal clearing, cursor positioning, and text rendering
3. **Input Handler**: Processes keyboard input with proper signal handling
4. **Font System**: Custom block-style digital font for large number display

### Display Algorithm

1. Detect terminal dimensions using `golang.org/x/term`
2. Calculate center position: `width/2`, `height/2`
3. Render each line using Unicode-aware width calculation
4. Position cursor precisely for perfect centering

### Blinking Implementation

True blinking is achieved by:
- Alternating between showing and hiding the display
- 200ms intervals for 5 seconds total (25 cycles)
- Clear screen during "off" phases
- Maintain bottom instruction text visibility

## Development

### Project Structure

```
commandline-timer/
├── main.go          # Main application code
├── go.mod          # Go module file
├── go.sum          # Go module checksums
├── Makefile        # Build automation
├── README.md       # This file
└── timer           # Built binary (after build)
```

### Building and Testing

```bash
# Build the project
make build

# Run tests
make test

# Clean up
make clean
```

## How It Works

1. **Terminal Setup**: Puts terminal in raw mode for direct input handling
2. **Size Detection**: Gets terminal dimensions for centering calculations
3. **Main Loop**: Updates display every 100ms for smooth countdown
4. **Input Handling**: Monitors for keyboard input without blocking
5. **Completion**: Triggers blinking alert when timer reaches zero
6. **Cleanup**: Restores terminal state on exit

## Troubleshooting

### Common Issues

**"Permission denied" when installing**
- Use `sudo make install` for system-wide installation
- Or install to user directory: `mkdir -p ~/bin && cp timer ~/bin/`

**"command not found" after installation**
- Ensure `/usr/local/bin` is in your PATH
- Add to shell profile: `export PATH="/usr/local/bin:$PATH"`

**Display issues**
- Terminal must support ANSI colors and UTF-8
- Try resizing terminal window
- Check terminal compatibility

**Controls not working**
- Ensure terminal is in focus
- Try Ctrl+C if other controls fail
- Check for conflicting key mappings

## License

This project is open source and available under the MIT License.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly on different terminal sizes
5. Submit a pull request

## Version History

- **v1.0.0**: Initial release with centered display and HH:MM:SS format
- **v1.1.0**: Added proper blinking behavior and smart controls
- **v1.2.0**: Improved Unicode support and centering accuracy

## Support

If you encounter any issues, please check:
1. Go version (1.21+ required)
2. Terminal compatibility (must support ANSI codes)
3. File permissions
4. PATH configuration