package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Terminal struct {
	oldState *term.State
}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (t *Terminal) GetSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Fallback to reasonable defaults
		return 80, 24
	}
	return width, height
}

func (t *Terminal) SetRawMode() error {
	// Check if stdin is a terminal
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return fmt.Errorf("stdin is not a terminal")
	}
	
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("failed to set terminal to raw mode: %v", err)
	}
	t.oldState = oldState
	return nil
}

func (t *Terminal) Restore() {
	if t.oldState != nil {
		term.Restore(int(os.Stdin.Fd()), t.oldState)
	}
}

func (t *Terminal) HideCursor() {
	fmt.Print("\033[?25l")
}

func (t *Terminal) ShowCursor() {
	fmt.Print("\033[?25h")
}

func (t *Terminal) ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

func (t *Terminal) MoveCursorTo(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func (t *Terminal) SetColor(color string) {
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
		"dark":     "\033[90m",
	}
	if code, exists := colors[color]; exists {
		fmt.Print(code)
	}
}

func (t *Terminal) CenterText(text string, width int) string {
	textWidth := len(text)
	if textWidth >= width {
		return text
	}
	padding := (width - textWidth) / 2
	return fmt.Sprintf("%*s%s", padding, "", text)
}

func (t *Terminal) Flush() {
	// Ensure all output is written to the terminal
	// This is a placeholder for any necessary flushing operations
}