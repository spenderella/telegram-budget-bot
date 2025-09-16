package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	BotToken       string `json:"-"`               // из env
	PollingTimeout int    `json:"polling_timeout"` // из файла
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	config.BotToken = os.Getenv("BOT_TOKEN")
	if config.BotToken == "" {
		return nil, fmt.Errorf("BOT_TOKEN environment variable is not set")
	}

	if config.PollingTimeout <= 0 {
		return nil, fmt.Errorf("polling_timeout must be positive")
	}

	return &config, nil
}
