package display

import (
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/tomcat/commandline-timer/pkg/timer"
	"github.com/tomcat/commandline-timer/pkg/terminal"
)

type Display struct {
	terminal *terminal.Terminal
}

func NewDisplay(t *terminal.Terminal) *Display {
	return &Display{
		terminal: t,
	}
}

func (d *Display) DrawTimer(tmr interface{}) {
	d.terminal.ClearScreen()
	
	width, height := d.terminal.GetSize()
	
	var remaining time.Duration
	var timerText string
	var stateText string
	var isTomato bool
	
	// Handle both regular Timer and TomatoTimer
	switch t := tmr.(type) {
	case *timer.Timer:
		remaining = t.Remaining()
		isTomato = false
	case *timer.TomatoTimer:
		remaining = t.Remaining()
		isTomato = true
		stateText = t.State()
	}
	
	totalSeconds := int(remaining.Seconds())
	
	// Always display in HH:MM:SS format
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	timerText = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	
	// Calculate positions
	centerY := height / 2
	centerX := width / 2
	
	// Create large text representation
	largeText := d.createLargeText(timerText)
	textLines := len(largeText)
	startY := centerY - textLines/2
	
	// Draw large timer with clean white text
	d.terminal.SetColor("white")
	d.terminal.SetColor("bold")
	for i, line := range largeText {
		// Calculate the actual width of the line (including Unicode characters)
		lineWidth := runewidth.StringWidth(line)
		// Center the line properly by positioning the start column
		startCol := centerX - lineWidth/2
		d.terminal.MoveCursorTo(startY+i, startCol)
		fmt.Print(line)
	}
	
	// Draw state indicator for tomato mode
	if isTomato {
		stateY := startY - 2
		var stateColor, stateLabel string
		
		switch stateText {
		case "focus":
			stateColor = "red"
			stateLabel = "🍅 FOCUS TIME"
		case "break":
			stateColor = "green"
			stateLabel = "☕ BREAK TIME"
		}
		
		d.terminal.MoveCursorTo(stateY, 0)
		d.terminal.SetColor(stateColor)
		d.terminal.SetColor("bold")
		fmt.Print(d.terminal.CenterText(stateLabel, width))
	}
	
	// Draw small instructions at bottom
	d.terminal.MoveCursorTo(height-1, 0)
	d.terminal.SetColor("reset")
	if isTomato {
		var instruction string
		if stateText == "focus" {
			instruction = "Press Ctrl+Q or Ctrl+C to exit"
		} else {
			instruction = "Press Ctrl+Q or Ctrl+C to exit"
		}
		fmt.Print(d.terminal.CenterText(instruction, width))
	} else {
		fmt.Print(d.terminal.CenterText("Press Ctrl+Q or Ctrl+C to exit", width))
	}
	
	// Reset colors
	d.terminal.SetColor("reset")
}

func (d *Display) FlashZero(keyChan chan byte) {
	width, height := d.terminal.GetSize()
	centerY := height / 2
	centerX := width / 2
	
	// Create large "00:00:00"
	largeZero := d.createLargeText("00:00:00")
	textLines := len(largeZero)
	startY := centerY - textLines/2
	
	// Flash for 5 seconds (25 intervals of 200ms)
	for i := 0; i < 25; i++ {
		select {
		case key := <-keyChan:
			if key == 13 || key == 10 { // Enter key (CR or LF)
				return
			}
		default:
			// Continue with flashing
		}
		
		d.terminal.ClearScreen()
		
		// Only draw the text when it should be visible (on even intervals)
		if i%2 == 0 {
			// Flash on - bright bold white
			d.terminal.SetColor("white")
			d.terminal.SetColor("bold")
			
			// Draw large 00:00:00
			for j, line := range largeZero {
				// Calculate the actual width of the line (including Unicode characters)
				lineWidth := runewidth.StringWidth(line)
				// Center the line properly by positioning the start column
				startCol := centerX - lineWidth/2
				d.terminal.MoveCursorTo(startY+j, startCol)
				fmt.Print(line)
			}
		}
		// When i%2 == 1, we don't draw anything - this creates the blink effect
		// The screen is already cleared, so nothing will be visible
		
		// Always draw instructions
		d.terminal.MoveCursorTo(height-1, 0)
		d.terminal.SetColor("reset")
		fmt.Print(d.terminal.CenterText("Press Enter to stop flashing or Ctrl+C to exit", width))
		
		d.terminal.Flush()
		os.Stdout.Sync()
		time.Sleep(200 * time.Millisecond)
	}
	
	// Final reset (don't clear screen to allow Time's up! message to show)
	d.terminal.SetColor("reset")
	// Move cursor to bottom to make room for Time's up! message
	d.terminal.MoveCursorTo(height-2, 0)
}

func (d *Display) FlashTomatoFocusComplete(keyChan chan byte) {
	width, height := d.terminal.GetSize()
	centerY := height / 2
	centerX := width / 2
	
	// Create large "BREAK"
	largeBreak := d.createLargeText("BREAK")
	textLines := len(largeBreak)
	startY := centerY - textLines/2
	
	// Flash until Enter is pressed
	for {
		select {
		case key := <-keyChan:
			if key == 13 || key == 10 { // Enter key (CR or LF)
				return
			}
		default:
			// Continue flashing
		}
		
		d.terminal.ClearScreen()
		
		// Draw focus complete message
		d.terminal.MoveCursorTo(startY-3, 0)
		d.terminal.SetColor("red")
		d.terminal.SetColor("bold")
		fmt.Print(d.terminal.CenterText("🍅 FOCUS COMPLETE!", width))
		
		d.terminal.MoveCursorTo(startY-1, 0)
		d.terminal.SetColor("white")
		fmt.Print(d.terminal.CenterText("Press Enter to start break", width))
		
		// Draw BREAK text
		for i, line := range largeBreak {
			lineWidth := runewidth.StringWidth(line)
			startCol := centerX - lineWidth/2
			d.terminal.MoveCursorTo(startY+i, startCol)
			d.terminal.SetColor("green")
			d.terminal.SetColor("bold")
			fmt.Print(line)
		}
		
		d.terminal.MoveCursorTo(height-1, 0)
		d.terminal.SetColor("reset")
		fmt.Print(d.terminal.CenterText("Press Enter to continue or Ctrl+C to exit", width))
		
		d.terminal.Flush()
		os.Stdout.Sync()
		time.Sleep(300 * time.Millisecond)
	}
}

func (d *Display) FlashTomatoBreakComplete(keyChan chan byte) {
	width, height := d.terminal.GetSize()
	centerY := height / 2
	centerX := width / 2
	
	// Create large "DONE"
	largeDone := d.createLargeText("DONE")
	textLines := len(largeDone)
	startY := centerY - textLines/2
	
	// Flash for 5 seconds or until Enter is pressed
	for i := 0; i < 17; i++ { // 17 * 300ms = ~5 seconds
		select {
		case key := <-keyChan:
			if key == 13 || key == 10 { // Enter key (CR or LF)
				return
			}
		default:
			// Continue flashing
		}
		
		d.terminal.ClearScreen()
		
		// Only draw on even intervals for blinking effect
		if i%2 == 0 {
			// Draw break complete message
			d.terminal.MoveCursorTo(startY-3, 0)
			d.terminal.SetColor("green")
			d.terminal.SetColor("bold")
			fmt.Print(d.terminal.CenterText("☕ BREAK COMPLETE!", width))
			
			d.terminal.MoveCursorTo(startY-1, 0)
			d.terminal.SetColor("white")
			fmt.Print(d.terminal.CenterText("Pomodoro cycle finished!", width))
			
			// Draw DONE text
			for j, line := range largeDone {
				lineWidth := runewidth.StringWidth(line)
				startCol := centerX - lineWidth/2
				d.terminal.MoveCursorTo(startY+j, startCol)
				d.terminal.SetColor("yellow")
				d.terminal.SetColor("bold")
				fmt.Print(line)
			}
		}
		
		d.terminal.MoveCursorTo(height-1, 0)
		d.terminal.SetColor("reset")
		fmt.Print(d.terminal.CenterText("Press Enter to exit or Ctrl+C to exit", width))
		
		d.terminal.Flush()
		os.Stdout.Sync()
		time.Sleep(300 * time.Millisecond)
	}
}

func (d *Display) Flush() {
	// Ensure all output is written to the terminal
	// This is a placeholder for any necessary flushing operations
}

func (d *Display) createLargeText(text string) []string {
	// Clean, readable digital font
	font := map[rune][]string{
		'0': {
			"█████",
			"█   █",
			"█   █",
			"█   █",
			"█████",
		},
		'1': {
			"    █",
			"    █",
			"    █",
			"    █",
			"    █",
		},
		'2': {
			"█████",
			"    █",
			"█████",
			"█    ",
			"█████",
		},
		'3': {
			"█████",
			"    █",
			"█████",
			"    █",
			"█████",
		},
		'4': {
			"█   █",
			"█   █",
			"█████",
			"    █",
			"    █",
		},
		'5': {
			"█████",
			"█    ",
			"█████",
			"    █",
			"█████",
		},
		'6': {
			"█████",
			"█    ",
			"█████",
			"█   █",
			"█████",
		},
		'7': {
			"█████",
			"    █",
			"    █",
			"    █",
			"    █",
		},
		'8': {
			"█████",
			"█   █",
			"█████",
			"█   █",
			"█████",
		},
		'9': {
			"█████",
			"█   █",
			"█████",
			"    █",
			"█████",
		},
		':': {
			" ",
			"█",
			" ",
			"█",
			" ",
		},
		// Letters for tomato mode
		'B': {
			"█████",
			"█   █",
			"█████",
			"█   █",
			"█████",
		},
		'R': {
			"█████",
			"█   █",
			"█████",
			"█ █  ",
			"█  █ ",
		},
		'E': {
			"█████",
			"█    ",
			"█████",
			"█    ",
			"█████",
		},
		'A': {
			"█████",
			"█   █",
			"█████",
			"█   █",
			"█   █",
		},
		'K': {
			"█   █",
			"█ █  ",
			"███  ",
			"█ █  ",
			"█   █",
		},
		'D': {
			"████ ",
			"█   █",
			"█   █",
			"█   █",
			"████ ",
		},
		'O': {
			"█████",
			"█   █",
			"█   █",
			"█   █",
			"█████",
		},
		'N': {
			"█   █",
			"██  █",
			"█ █ █",
			"█  ██",
			"█   █",
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
				line += "      " // Space for unknown characters
			}
		}
		result = append(result, line)
	}
	
	return result
}