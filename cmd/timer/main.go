package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tomcat/commandline-timer/pkg/config"
	"github.com/tomcat/commandline-timer/pkg/display"
	"github.com/tomcat/commandline-timer/pkg/timer"
	"github.com/tomcat/commandline-timer/pkg/terminal"
)

func run() error {
	// Parse configuration
	config, err := config.ParseConfig()
	if err != nil {
		return err
	}

	// Initialize terminal
	term := terminal.NewTerminal()
	if err := term.SetRawMode(); err != nil {
		return err
	}
	defer func() {
		term.Restore()
		term.ShowCursor()
		term.SetColor("reset")
		// Don't clear screen to allow final messages to show
	}()

	term.HideCursor()

	// Initialize timer and display
	var tmr interface{}
	if config.IsTomato {
		tmr = timer.NewTomatoTimer()
	} else {
		tmr = timer.NewTimer(config.Duration)
	}
	display := display.NewDisplay(term)

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
			if config.IsTomato {
				tmr.(*timer.TomatoTimer).Stop()
			} else {
				tmr.(*timer.Timer).Stop()
			}
			return nil

		case key := <-keyChan:
			if key == 17 { // Ctrl+Q (ASCII code for DC1)
				if config.IsTomato {
					tmr.(*timer.TomatoTimer).Stop()
				} else {
					tmr.(*timer.Timer).Stop()
				}
				return nil
			}

		case <-ticker.C:
			if config.IsTomato {
				tomatoTimer := tmr.(*timer.TomatoTimer)
				
				if tomatoTimer.IsExpired() {
					display.DrawTimer(tomatoTimer)
					
					if tomatoTimer.State() == "focus" {
						// Focus time completed, wait for Enter to start break
						display.FlashTomatoFocusComplete(keyChan)
						tomatoTimer.StartBreak()
						fmt.Println("\nFocus time completed! Starting break...")
					} else if tomatoTimer.State() == "break" {
						// Break time completed
						display.FlashTomatoBreakComplete(keyChan)
						fmt.Println("\nBreak time completed! Pomodoro cycle finished.")
						return nil
					}
				}
				
				display.DrawTimer(tomatoTimer)
			} else {
				regularTimer := tmr.(*timer.Timer)
				
				if regularTimer.IsExpired() {
					display.DrawTimer(regularTimer)
					display.FlashZero(keyChan)
					fmt.Println("\nTime's up!")
					return nil
				}
				
				display.DrawTimer(regularTimer)
			}
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}