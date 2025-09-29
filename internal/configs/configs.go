package configs

import (
	"encoding/json"
	"os"

	"telegram-finance-bot/internal/errors"
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
		return nil, errors.ErrFailedToOpenConfig(err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, errors.ErrFailedToDecodeConfig(err)
	}

	config.BotToken = os.Getenv("BOT_TOKEN")
	if config.BotToken == "" {
		return nil, errors.ErrBotTokenMissing
	}

	if config.PollingTimeout <= 0 {
		return nil, errors.ErrInvalidPollingTimeout
	}

	if config.GetExpenseLimit <= 0 {
		return nil, errors.ErrInvalidGetExpenseLimit
	}

	if config.GetExpenseLimit > MaxGetExpensesLimit {
		return nil, errors.ErrGetExpenseLimitTooLarge(config.GetExpenseLimit, MaxGetExpensesLimit)
	}

	return &config, nil
}
