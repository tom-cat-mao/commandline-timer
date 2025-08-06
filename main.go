package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mattn/go-runewidth"
	"golang.org/x/term"
)

type Timer struct {
	duration time.Duration
	started  time.Time
	running  bool
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		duration: duration,
		started:  time.Now(),
		running:  true,
	}
}

func (t *Timer) Remaining() time.Duration {
	if !t.running {
		return 0
	}
	elapsed := time.Since(t.started)
	remaining := t.duration - elapsed
	if remaining <= 0 {
		return 0
	}
	return remaining
}

func (t *Timer) IsExpired() bool {
	return t.Remaining() == 0
}

func (t *Timer) Stop() {
	t.running = false
}

func getTerminalSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Fallback to reasonable defaults
		return 80, 24
	}
	return width, height
}

func centerText(text string, width int) string {
	textWidth := runewidth.StringWidth(text)
	if textWidth >= width {
		return text
	}
	padding := (width - textWidth) / 2
	return fmt.Sprintf("%*s%s", padding, "", text)
}

func moveCursorTo(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func setColor(color string) {
	colors := map[string]string{
		"reset":    "\033[0m",
		"red":      "\033[31m",
		"green":    "\033[32m",
		"yellow":   "\033[33m",
		"blue":     "\033[34m",
		"magenta":  "\033[35m",
		"cyan":     "\033[36m",
		"white":    "\033[37m",
		"bold":     "\033[1m",
		"bg_red":   "\033[41m",
		"bg_green": "\033[42m",
	}
	if code, exists := colors[color]; exists {
		fmt.Print(code)
	}
}

func drawTimer(timer *Timer) {
	clearScreen()
	
	width, height := getTerminalSize()
	
	remaining := timer.Remaining()
	totalSeconds := int(remaining.Seconds())
	
	// Calculate positions
	centerY := height / 2
	centerX := width / 2
	
	// Format time based on duration
	var timerText string
	if timer.duration >= time.Hour {
		hours := totalSeconds / 3600
		minutes := (totalSeconds % 3600) / 60
		seconds := totalSeconds % 60
		timerText = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	} else if timer.duration >= time.Minute {
		minutes := totalSeconds / 60
		seconds := totalSeconds % 60
		timerText = fmt.Sprintf("%02d:%02d", minutes, seconds)
	} else {
		timerText = fmt.Sprintf("%d", totalSeconds)
	}
	
	// Create large text representation
	largeText := createLargeText(timerText)
	textLines := len(largeText)
	startY := centerY - textLines/2
	
	// Draw large timer
	setColor("white")
	setColor("bold")
	for i, line := range largeText {
		moveCursorTo(startY+i, centerX-len(line)/2)
		fmt.Print(line)
	}
	
	// Draw small instructions at bottom
	moveCursorTo(height-1, 0)
	setColor("reset")
	setColor("cyan")
	fmt.Print(centerText("Press Enter or Ctrl+C to exit", width))
	
	// Reset colors
	setColor("reset")
	os.Stdout.Sync()
}

func createLargeText(text string) []string {
	// Large digit font (3x5 for numbers, 1x5 for colon)
	font := map[rune][]string{
		'0': {
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		'1': {
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		},
		'2': {
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		'3': {
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		'4': {
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		'5': {
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		'6': {
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		},
		'7': {
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		},
		'8': {
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		'9': {
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		},
		':': {
			"  ",
			"██",
			"  ",
			"██",
			"  ",
		},
		' ': {
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
	}
	
	// Convert each character to large text
	var result []string
	for row := 0; row < 5; row++ {
		var line string
		for _, char := range text {
			if charLines, exists := font[char]; exists {
				line += charLines[row] + " "
			} else {
				line += "    " // Space for unknown characters
			}
		}
		result = append(result, line)
	}
	
	return result
}

func flashZero() {
	width, height := getTerminalSize()
	centerY := height / 2
	centerX := width / 2
	
	// Create large "0"
	largeZero := createLargeText("0")
	textLines := len(largeZero)
	startY := centerY - textLines/2
	
	for i := 0; i < 25; i++ { // 5 seconds at 200ms intervals
		if i%2 == 0 {
			// Flash on - red background
			setColor("bg_red")
			setColor("white")
			setColor("bold")
		} else {
			// Flash off - normal
			setColor("reset")
			setColor("white")
			setColor("bold")
		}
		
		clearScreen()
		
		// Draw large zero
		for j, line := range largeZero {
			moveCursorTo(startY+j, centerX-len(line)/2)
			fmt.Print(line)
		}
		
		os.Stdout.Sync()
		time.Sleep(200 * time.Millisecond)
	}
	
	setColor("reset")
}

func run() error {
	flag.Parse()
	
	if flag.NArg() == 0 {
		fmt.Println("Usage: timer <duration>")
		fmt.Println("Examples:")
		fmt.Println("  timer 30s    # 30 seconds")
		fmt.Println("  timer 5m     # 5 minutes")
		fmt.Println("  timer 1h30m  # 1 hour 30 minutes")
		os.Exit(1)
	}
	
	durationStr := flag.Arg(0)
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Printf("Invalid duration: %v\n", err)
		os.Exit(1)
	}
	
	if duration <= 0 {
		fmt.Println("Duration must be positive")
		os.Exit(1)
	}
	
	// Save terminal state
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("failed to set terminal to raw mode: %v", err)
	}
	defer func() {
		term.Restore(int(os.Stdin.Fd()), oldState)
		showCursor()
		setColor("reset")
		clearScreen()
	}()
	
	hideCursor()
	
	timer := NewTimer(duration)
	
	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Create ticker for updates
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	// Channel for key input
	keyChan := make(chan byte, 1)
	
	// Start key reader goroutine
	go func() {
		buf := make([]byte, 1)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				return
			}
			if n > 0 {
				keyChan <- buf[0]
			}
		}
	}()
	
	for {
		select {
		case <-sigChan:
			timer.Stop()
			return nil
			
		case key := <-keyChan:
			if key == 13 || key == 10 { // Enter key (CR or LF)
				timer.Stop()
				return nil
			}
			
		case <-ticker.C:
			if timer.IsExpired() {
				drawTimer(timer)
				flashZero()
				fmt.Println("\nTime's up!")
				return nil
			}
			
			drawTimer(timer)
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}