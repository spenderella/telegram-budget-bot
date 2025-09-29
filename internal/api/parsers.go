package api

import (
	"strconv"
	"strings"
	"telegram-finance-bot/internal/constants"
	"telegram-finance-bot/internal/errors"
)

func parseAddExpenseCommand(text string) (float64, string, error) {
	// "/add_expense 100 food" → ["add_expense", "100", "food"]
	parts := strings.Fields(text)

	if len(parts) < 3 {
		return 0, "", errors.ErrInvalidCommandFormat(constants.CmdAddExpense)
	}

	amount, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, "", errors.ErrInvalidAmountFormat
	}

	if amount <= 0 {
		return 0, "", errors.ErrAmountMustBePositive
	}

	category := strings.Join(parts[2:], " ")

	return amount, category, nil
}
