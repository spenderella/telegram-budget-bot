package api

import (
	"errors"
	"strconv"
	"strings"
)

func parseAddExpenseCommand(text string) (float64, string, error) {
	// "/add_expense 100 food" → ["add_expense", "100", "food"]
	parts := strings.Fields(text)

	if len(parts) < 3 {
		return 0, "", errors.New("invalid format. Use: /add_expense <amount> <category>")
	}

	amount, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, "", errors.New("invalid amount format")
	}

	if amount <= 0 {
		return 0, "", errors.New("amount must be positive")
	}

	category := strings.Join(parts[2:], " ")

	return amount, category, nil
}
