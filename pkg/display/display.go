package display

import (
	"fmt"

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

func (d *Display) DrawTimer(timer *timer.Timer) {
	d.terminal.ClearScreen()
	
	width, height := d.terminal.GetSize()
	
	remaining := timer.Remaining()
	totalSeconds := int(remaining.Seconds())
	
	// Always display in HH:MM:SS format
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	timerText := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	
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
	
	// Draw small instructions at bottom
	d.terminal.MoveCursorTo(height-1, 0)
	d.terminal.SetColor("reset")
	fmt.Print(d.terminal.CenterText("Press Ctrl+Q or Ctrl+C to exit", width))
	
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
		
		// Always draw instructions
		d.terminal.MoveCursorTo(height-1, 0)
		d.terminal.SetColor("reset")
		fmt.Print(d.terminal.CenterText("Press Enter to stop flashing or Ctrl+C to exit", width))
		
		d.terminal.Flush()
	}
	
	// Final clear and reset
	d.terminal.SetColor("reset")
	d.terminal.ClearScreen()
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