package errors

import (
	"errors"
	"fmt"

	"telegram-finance-bot/internal/constants"
)

func ErrInvalidCommandFormat(commandType string) error {
	var format string
	var examples []string

	switch commandType {
	case constants.CmdAddExpense:
		format = constants.AddExpenseFormat
		examples = []string{constants.AddExpenseExample}

	default:
		return fmt.Errorf("unknown command: %s", commandType)
	}

	msg := fmt.Sprintf("Invalid format. Use: %s\n\nExamples:\n", format)
	for _, example := range examples {
		msg += fmt.Sprintf("• %s\n", example)
	}

	return errors.New(msg)
}

var (
	ErrInvalidAmountFormat  = errors.New("invalid amount format")
	ErrAmountMustBePositive = errors.New("amount must be positive")
)
