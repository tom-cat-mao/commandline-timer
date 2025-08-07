package timer

import (
	"testing"
	"time"
)

func TestRegularTimer(t *testing.T) {
	timer := NewTimer(1 * time.Second)
	
	// Test initial state
	if timer.Remaining() > 1*time.Second || timer.Remaining() <= 0 {
		t.Errorf("Expected remaining time ~1s, got %v", timer.Remaining())
	}
	
	if timer.IsExpired() {
		t.Error("Expected timer to not be expired initially")
	}
	
	// Test after duration
	time.Sleep(1100 * time.Millisecond)
	if !timer.IsExpired() {
		t.Error("Expected timer to be expired after 1.1s")
	}
}

func TestTomatoTimer(t *testing.T) {
	tomatoTimer := NewTomatoTimer()
	
	// Test initial state
	if tomatoTimer.State() != "focus" {
		t.Errorf("Expected initial state 'focus', got %v", tomatoTimer.State())
	}
	
	if tomatoTimer.Remaining() > 25*time.Minute || tomatoTimer.Remaining() <= 0 {
		t.Errorf("Expected remaining time ~25m, got %v", tomatoTimer.Remaining())
	}
	
	// Test state transition
	tomatoTimer.StartBreak()
	if tomatoTimer.State() != "break" {
		t.Errorf("Expected state 'break' after StartBreak, got %v", tomatoTimer.State())
	}
	
	if tomatoTimer.Remaining() > 5*time.Minute || tomatoTimer.Remaining() <= 0 {
		t.Errorf("Expected remaining time ~5m, got %v", tomatoTimer.Remaining())
	}
	
	// Test reset
	tomatoTimer.Reset()
	if tomatoTimer.State() != "focus" {
		t.Errorf("Expected state 'focus' after Reset, got %v", tomatoTimer.State())
	}
}

func TestTimerReset(t *testing.T) {
	timer := NewTimer(1 * time.Second)
	time.Sleep(1100 * time.Millisecond)
	
	if !timer.IsExpired() {
		t.Error("Expected timer to be expired")
	}
	
	// Reset timer
	timer.Reset(2 * time.Second)
	if timer.IsExpired() {
		t.Error("Expected timer to not be expired after reset")
	}
	
	if timer.Remaining() > 2*time.Second || timer.Remaining() <= 0 {
		t.Errorf("Expected remaining time ~2s, got %v", timer.Remaining())
	}
}