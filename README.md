# Command Line Timer

A cross-platform command line countdown timer written in Go. It displays a centered countdown timer in the terminal with a progress bar and visual feedback when time expires.

## Features

- 🖥️ **Cross-platform**: Works on Linux, macOS, and other Unix-like systems
- 🎯 **Centered display**: Timer is automatically centered in the terminal
- 📊 **Progress bar**: Visual progress indicator
- 🎨 **Colorful interface**: Uses colors for better visibility
- ⏰ **Multiple time formats**: Supports seconds, minutes, and hours
- 🛑 **Graceful exit**: Clean shutdown with Ctrl+C
- 🎭 **Completion alert**: Visual flash when timer expires

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

### Build Script

For easy building, use the provided build script:
```bash
# Build for current platform
./build.sh

# Build for all platforms
./build.sh --all
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
   sudo cp timer /usr/local/bin/
   ```

## Usage

### Basic Usage

```bash
timer 30s    # 30 seconds
timer 5m     # 5 minutes
timer 1h30m  # 1 hour 30 minutes
timer 2h45m30s  # 2 hours, 45 minutes, 30 seconds
```

### Examples

```bash
# 10-second countdown
timer 10s

# 5-minute countdown
timer 5m

# 1-hour countdown
timer 1h

# Mixed format
timer 1h30m15s
```

### Controls

- **Ctrl+C**: Exit the timer
- Timer automatically exits when time is up

## Building

### Build for Current Platform

```bash
make build
```

### Build for Specific Platforms

```bash
make build-linux      # Linux (x86_64)
make build-macos      # macOS (Intel)
make build-macos-arm64 # macOS (ARM64)
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
  - `github.com/mattn/go-runewidth`
  - `github.com/nsf/termbox-go`

## Development

### Project Structure

```
commandline-timer/
├── main.go          # Main application code
├── go.mod          # Go module file
├── go.sum          # Go module checksums
├── Makefile        # Build automation
├── build.sh        # Build script
├── README.md       # This file
└── timer           # Built binary (after build)
```

### Building and Testing

```bash
# Download dependencies
make deps

# Build the project
make build

# Run tests (if any)
make test

# Clean up
make clean
```

## How It Works

1. **Terminal Detection**: Uses `termbox-go` to detect terminal size and handle input
2. **Centered Display**: Calculates center position based on terminal dimensions
3. **Progress Bar**: Shows visual progress of the countdown
4. **Real-time Updates**: Updates display every 100ms for smooth animation
5. **Cross-platform**: Works on both Linux and macOS systems

## License

This project is open source and available under the MIT License.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## Troubleshooting

### Common Issues

**"Permission denied" when installing**
- Use `sudo make install` to install system-wide
- Or install to user directory: `mkdir -p ~/bin && cp timer ~/bin/`

**"command not found" after installation**
- Ensure `/usr/local/bin` is in your PATH
- Add to your shell profile: `export PATH="/usr/local/bin:$PATH"`

**Display issues**
- Make sure your terminal supports ANSI colors
- Try resizing your terminal window
- Check that your terminal supports UTF-8 characters

### Support

If you encounter any issues, please check:
1. Go version (1.21+ required)
2. Terminal compatibility
3. File permissions
4. PATH configuration