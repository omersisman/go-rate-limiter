package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

var (
	config Config
	mu     sync.RWMutex
)

type Config struct {
	RateLimits map[string]RateLimitConfig `yaml:"rate_limits"`
}

type RateLimitConfig struct {
	Limit    int `yaml:"limit"`
	Duration int `yaml:"duration"`
}

func LoadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
}

func GetRateLimit(endpoint string) (RateLimitConfig, bool) {
	rl, exists := config.RateLimits[endpoint]
	return rl, exists
}
