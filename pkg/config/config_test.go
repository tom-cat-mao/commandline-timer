package config

import (
	"os"
	"testing"
	"time"
)

func TestParseRegularTimer(t *testing.T) {
	// Save original os.Args and restore after test
	originalArgs := []string{"timer", "5m"}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	
	os.Args = originalArgs
	
	config, err := ParseConfig()
	if err != nil {
		t.Fatalf("ParseConfig failed: %v", err)
	}
	
	if config.IsTomato {
		t.Error("Expected IsTomato to be false for regular timer")
	}
	
	if config.Duration != 5*time.Minute {
		t.Errorf("Expected duration 5m, got %v", config.Duration)
	}
	
	if config.TomatoState != "" {
		t.Errorf("Expected empty TomatoState, got %v", config.TomatoState)
	}
}

func TestParseTomatoTimer(t *testing.T) {
	// Save original os.Args and restore after test
	originalArgs := []string{"timer", "tomato"}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	
	os.Args = originalArgs
	
	config, err := ParseConfig()
	if err != nil {
		t.Fatalf("ParseConfig failed: %v", err)
	}
	
	if !config.IsTomato {
		t.Error("Expected IsTomato to be true for tomato timer")
	}
	
	if config.Duration != 25*time.Minute {
		t.Errorf("Expected duration 25m, got %v", config.Duration)
	}
	
	if config.TomatoState != "focus" {
		t.Errorf("Expected TomatoState 'focus', got %v", config.TomatoState)
	}
}