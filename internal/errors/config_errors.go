package errors

import (
	"errors"
	"fmt"
)

// Base errors
var (
	ErrBotTokenMissing        = errors.New("BOT_TOKEN environment variable is not set")
	ErrInvalidPollingTimeout  = errors.New("polling_timeout must be positive")
	ErrInvalidGetExpenseLimit = errors.New("get_expense_limit must be positive")
)

func ErrFailedToOpenConfig(err error) error {
	return fmt.Errorf("failed to open config file: %w", err)
}

func ErrFailedToDecodeConfig(err error) error {
	return fmt.Errorf("failed to decode config: %w", err)
}

func ErrConfigValidation(field string, value interface{}) error {
	return fmt.Errorf("%s must be positive, got: %v", field, value)
}

func ErrGetExpenseLimitTooLarge(limit, max int) error {
	return fmt.Errorf("get_expense_limit must be less than %d, got: %d", max, limit)
}

func ErrFailedToCreateBot(err error) error {
	return fmt.Errorf("failed to create bot: %w", err)
}
