package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	MaxGetExpensesLimit = 500
)

type Config struct {
	BotToken        string `json:"-"`                 // from env
	PollingTimeout  int    `json:"polling_timeout"`   // from file
	GetExpenseLimit int    `json:"get_expense_limit"` // from file

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

	if config.GetExpenseLimit <= 0 {
		return nil, fmt.Errorf("get_expense_limit must be positive")
	}

	if config.GetExpenseLimit > MaxGetExpensesLimit {
		return nil, fmt.Errorf("get_expense_limit must be less than %d", MaxGetExpensesLimit)
	}

	return &config, nil
}
