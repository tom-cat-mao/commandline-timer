package timer

import (
	"time"
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

func (t *Timer) Reset(duration time.Duration) {
	t.duration = duration
	t.started = time.Now()
	t.running = true
}

// TomatoTimer represents a pomodoro timer with focus and break states
type TomatoTimer struct {
	focusTimer *Timer
	breakTimer *Timer
	state      string // "focus", "break", "completed"
}

func NewTomatoTimer() *TomatoTimer {
	return &TomatoTimer{
		focusTimer: NewTimer(25 * time.Minute),
		breakTimer: NewTimer(5 * time.Minute),
		state:      "focus",
	}
}

func (tt *TomatoTimer) CurrentTimer() *Timer {
	if tt.state == "focus" {
		return tt.focusTimer
	} else if tt.state == "break" {
		return tt.breakTimer
	}
	return tt.focusTimer // fallback
}

func (tt *TomatoTimer) Remaining() time.Duration {
	return tt.CurrentTimer().Remaining()
}

func (tt *TomatoTimer) IsExpired() bool {
	return tt.CurrentTimer().IsExpired()
}

func (tt *TomatoTimer) Stop() {
	tt.focusTimer.Stop()
	tt.breakTimer.Stop()
}

func (tt *TomatoTimer) State() string {
	return tt.state
}

func (tt *TomatoTimer) StartBreak() {
	tt.state = "break"
	tt.breakTimer.Reset(5 * time.Minute)
}

func (tt *TomatoTimer) Reset() {
	tt.state = "focus"
	tt.focusTimer.Reset(25 * time.Minute)
	tt.breakTimer.Reset(5 * time.Minute)
}