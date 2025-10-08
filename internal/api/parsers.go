package api

import (
	"strconv"
	"strings"
	"telegram-finance-bot/internal/constants"
	"telegram-finance-bot/internal/errors"
	"time"
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

func (bot *BudgetBot) getPeriodDates(period string) (time.Time, time.Time) {
	now := time.Now().UTC()

	switch period {
	case "today":
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, time.UTC)
		return start, end

	case "week":
		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		daysToMonday := weekday - 1
		start := time.Date(now.Year(), now.Month(), now.Day()-daysToMonday, 0, 0, 0, 0, time.UTC)
		end := now
		return start, end

	case "month":
		start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		end := now
		return start, end

	default:
		return time.Time{}, now
	}
}
