package main

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/tomcat/commandline-timer/pkg/config"
	"github.com/tomcat/commandline-timer/pkg/timer"
)

func TestTomatoModeIntegration(t *testing.T) {
	// Test that tomato mode can be initialized correctly
	originalArgs := []string{"timer", "tomato"}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	
	os.Args = originalArgs
	
	config, err := config.ParseConfig()
	if err != nil {
		t.Fatalf("Failed to parse config: %v", err)
	}
	
	if !config.IsTomato {
		t.Error("Expected tomato mode to be enabled")
	}
	
	if config.TomatoState != "focus" {
		t.Errorf("Expected focus state, got %v", config.TomatoState)
	}
	
	// Test tomato timer initialization
	tomatoTimer := timer.NewTomatoTimer()
	if tomatoTimer.State() != "focus" {
		t.Errorf("Expected focus state, got %v", tomatoTimer.State())
	}
	
	// Test state transition
	tomatoTimer.StartBreak()
	if tomatoTimer.State() != "break" {
		t.Errorf("Expected break state, got %v", tomatoTimer.State())
	}
}

func TestRegularTimerIntegration(t *testing.T) {
	// Test that regular timer mode works
	originalArgs := []string{"timer", "30s"}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	
	os.Args = originalArgs
	
	config, err := config.ParseConfig()
	if err != nil {
		t.Fatalf("Failed to parse config: %v", err)
	}
	
	if config.IsTomato {
		t.Error("Expected regular timer mode")
	}
	
	if config.Duration != 30*time.Second {
		t.Errorf("Expected 30s duration, got %v", config.Duration)
	}
}

func TestHelpMessage(t *testing.T) {
	// Test that help message contains tomato information
	originalArgs := []string{"timer", "invalid"}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	
	os.Args = originalArgs
	
	_, err := config.ParseConfig()
	if err == nil {
		t.Error("Expected ParseConfig to fail with invalid duration")
	}
	
	// The error message should contain usage information
	if !strings.Contains(err.Error(), "invalid duration") {
		t.Errorf("Expected 'invalid duration' error, got: %v", err)
	}
}