package config

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Duration    time.Duration
	IsTomato    bool
	TomatoState string // "focus", "break", "completed"
}

func ParseConfig() (*Config, error) {
	flag.Parse()
	
	if flag.NArg() == 0 {
		fmt.Println("Usage: timer <duration> OR timer tomato")
		fmt.Println("Examples:")
		fmt.Println("  timer 30s    # 30 seconds")
		fmt.Println("  timer 5m     # 5 minutes")
		fmt.Println("  timer 1h30m  # 1 hour 30 minutes")
		fmt.Println("  timer tomato # 25 min focus + 5 min break")
		os.Exit(1)
	}
	
	firstArg := flag.Arg(0)
	
	// Check if it's tomato mode
	if firstArg == "tomato" {
		return &Config{
			Duration:    25 * time.Minute, // 25 minutes focus time
			IsTomato:    true,
			TomatoState: "focus",
		}, nil
	}
	
	// Regular timer mode
	duration, err := time.ParseDuration(firstArg)
	if err != nil {
		return nil, fmt.Errorf("invalid duration: %v", err)
	}
	
	if duration <= 0 {
		return nil, fmt.Errorf("duration must be positive")
	}
	
	return &Config{
		Duration:    duration,
		IsTomato:    false,
		TomatoState: "",
	}, nil
}