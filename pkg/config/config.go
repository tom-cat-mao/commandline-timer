package config

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Duration time.Duration
}

func ParseConfig() (*Config, error) {
	flag.Parse()
	
	if flag.NArg() == 0 {
		fmt.Println("Usage: timer <duration>")
		fmt.Println("Examples:")
		fmt.Println("  timer 30s    # 30 seconds")
		fmt.Println("  timer 5m     # 5 minutes")
		fmt.Println("  timer 1h30m  # 1 hour 30 minutes")
		os.Exit(1)
	}
	
	durationStr := flag.Arg(0)
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return nil, fmt.Errorf("invalid duration: %v", err)
	}
	
	if duration <= 0 {
		return nil, fmt.Errorf("duration must be positive")
	}
	
	return &Config{
		Duration: duration,
	}, nil
}