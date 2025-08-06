package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
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
	w, h := termbox.Size()
	if w == 0 || h == 0 {
		w, h = 80, 24
	}
	return w, h
}

func centerText(text string, width int) string {
	textWidth := runewidth.StringWidth(text)
	if textWidth >= width {
		return text
	}
	padding := (width - textWidth) / 2
	return fmt.Sprintf("%*s%s", padding, "", text)
}

func drawTimer(timer *Timer) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	
	width, height := getTerminalSize()
	
	remaining := timer.Remaining()
	seconds := int(remaining.Seconds())
	minutes := seconds / 60
	seconds = seconds % 60
	hours := minutes / 60
	minutes = minutes % 60
	
	timerText := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	
	// Center the timer
	y := height / 2
	x := 0
	
	// Draw timer text
	for i, r := range timerText {
		termbox.SetCell(x+i, y, r, termbox.ColorWhite, termbox.ColorDefault)
	}
	
	// Draw title
	title := "COUNTDOWN TIMER"
	titleX := (width - runewidth.StringWidth(title)) / 2
	for i, r := range title {
		termbox.SetCell(titleX+i, y-2, r, termbox.ColorCyan, termbox.ColorDefault)
	}
	
	// Draw progress bar
	progress := 1.0 - float64(remaining.Nanoseconds())/float64(timer.duration.Nanoseconds())
	barWidth := width / 2
	barX := (width - barWidth) / 2
	
	for i := 0; i < barWidth; i++ {
		if float64(i) < float64(barWidth)*progress {
			termbox.SetCell(barX+i, y+2, '=', termbox.ColorGreen, termbox.ColorDefault)
		} else {
			termbox.SetCell(barX+i, y+2, '-', termbox.ColorDarkGray, termbox.ColorDefault)
		}
	}
	
	// Draw instructions
	instructions := "Press Ctrl+C to exit"
	instX := (width - runewidth.StringWidth(instructions)) / 2
	for i, r := range instructions {
		termbox.SetCell(instX+i, height-2, r, termbox.ColorYellow, termbox.ColorDefault)
	}
	
	termbox.Flush()
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
	
	err = termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()
	
	termbox.SetInputMode(termbox.InputEsc)
	
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
				// Flash the screen when timer expires
				for i := 0; i < 5; i++ {
					termbox.Clear(termbox.ColorRed, termbox.ColorRed)
					termbox.Flush()
					time.Sleep(200 * time.Millisecond)
					termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
					termbox.Flush()
					time.Sleep(200 * time.Millisecond)
				}
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