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
	seconds := int(remaining.Seconds())
	minutes := seconds / 60
	seconds = seconds % 60
	hours := minutes / 60
	minutes = minutes % 60
	
	timerText := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	
	// Calculate positions
	centerY := height / 2
	titleY := centerY - 2
	timerY := centerY
	progressY := centerY + 2
	instructionsY := height - 2
	
	// Draw title
	moveCursorTo(titleY, 0)
	setColor("cyan")
	setColor("bold")
	fmt.Print(centerText("COUNTDOWN TIMER", width))
	
	// Draw timer
	moveCursorTo(timerY, 0)
	setColor("white")
	setColor("bold")
	fmt.Print(centerText(timerText, width))
	
	// Draw progress bar
	progress := 1.0 - float64(remaining.Nanoseconds())/float64(timer.duration.Nanoseconds())
	barWidth := width / 2
	barX := (width - barWidth) / 2
	
	moveCursorTo(progressY, barX)
	setColor("green")
	
	for i := 0; i < barWidth; i++ {
		if float64(i) < float64(barWidth)*progress {
			fmt.Print("=")
		} else {
			setColor("reset")
			fmt.Print("-")
			setColor("green")
		}
	}
	
	// Draw instructions
	moveCursorTo(instructionsY, 0)
	setColor("yellow")
	fmt.Print(centerText("Press Ctrl+C to exit", width))
	
	// Reset colors and move cursor to bottom
	setColor("reset")
	moveCursorTo(height, 0)
	
	os.Stdout.Sync()
}

func flashScreen() {
	for i := 0; i < 3; i++ {
		setColor("bg_red")
		clearScreen()
		os.Stdout.Sync()
		time.Sleep(200 * time.Millisecond)
		
		setColor("reset")
		clearScreen()
		os.Stdout.Sync()
		time.Sleep(200 * time.Millisecond)
	}
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
	
	// Handle Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Create ticker for updates
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-sigChan:
			timer.Stop()
			return nil
			
		case <-ticker.C:
			if timer.IsExpired() {
				drawTimer(timer)
				flashScreen()
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