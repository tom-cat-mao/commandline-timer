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