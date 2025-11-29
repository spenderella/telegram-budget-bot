package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"telegram-finance-bot/internal/errors"
)

func Test_parseAddExpenseCommand(t *testing.T) {
	t.Run("valid input with single word category", func(t *testing.T) {
		command := "/add_expense 100 food"

		amount, category, err := parseAddExpenseCommand(command)

		require.NoError(t, err)
		assert.Equal(t, 100.0, amount)
		assert.Equal(t, "food", category)
	})

	t.Run("valid input with multi-word category", func(t *testing.T) {
		command := "/add_expense 100 public transport"

		amount, category, err := parseAddExpenseCommand(command)

		require.NoError(t, err)
		assert.Equal(t, 100.0, amount)
		assert.Equal(t, "public transport", category)
	})

	t.Run("invalid format - missing category", func(t *testing.T) {
		command := "/add_expense 100"

		_, _, err := parseAddExpenseCommand(command)

		require.Error(t, err)
		// Check that it returns ErrInvalidCommandFormat
		assert.Contains(t, err.Error(), "Invalid format")
	})

	t.Run("invalid amount format", func(t *testing.T) {
		command := "/add_expense abc food"

		_, _, err := parseAddExpenseCommand(command)

		require.Error(t, err)
		// Check that it's specifically the invalid amount error
		assert.Equal(t, errors.ErrInvalidAmountFormat, err)
	})

	t.Run("negative amount", func(t *testing.T) {
		command := "/add_expense -50 food"

		_, _, err := parseAddExpenseCommand(command)

		require.Error(t, err)
		// Check that it's the "must be positive" error
		assert.Equal(t, errors.ErrAmountMustBePositive, err)
	})

	t.Run("zero amount", func(t *testing.T) {
		command := "/add_expense 0 food"

		_, _, err := parseAddExpenseCommand(command)

		require.Error(t, err)
		// Check that it's the "must be positive" error
		assert.Equal(t, errors.ErrAmountMustBePositive, err)
	})
}

func Test_getPeriodDates(t *testing.T) {
	// Create a BudgetBot instance for testing
	bot := &BudgetBot{}

	t.Run("today", func(t *testing.T) {
		// ACT
		start, end := bot.getPeriodDates("today")

		// ASSERT
		// Start should be at 00:00:00
		assert.Equal(t, 0, start.Hour())
		assert.Equal(t, 0, start.Minute())
		assert.Equal(t, 0, start.Second())

		// End should be at 23:59:59
		assert.Equal(t, 23, end.Hour())
		assert.Equal(t, 59, end.Minute())
		assert.Equal(t, 59, end.Second())

		// Both should be today
		now := time.Now().UTC()
		assert.Equal(t, now.Year(), start.Year())
		assert.Equal(t, now.Month(), start.Month())
		assert.Equal(t, now.Day(), start.Day())
		assert.Equal(t, now.Year(), end.Year())
		assert.Equal(t, now.Month(), end.Month())
		assert.Equal(t, now.Day(), end.Day())
	})

	t.Run("week", func(t *testing.T) {
		// ACT
		start, end := bot.getPeriodDates("week")

		// ASSERT
		// Start should be at Monday 00:00:00
		assert.Equal(t, time.Monday, start.Weekday())
		assert.Equal(t, 0, start.Hour())
		assert.Equal(t, 0, start.Minute())
		assert.Equal(t, 0, start.Second())

		// End should be approximately now
		now := time.Now().UTC()
		assert.WithinDuration(t, now, end, time.Second)

		// Start should be before or equal to end
		assert.True(t, start.Before(end) || start.Equal(end))
	})

	t.Run("month", func(t *testing.T) {
		// ACT
		start, end := bot.getPeriodDates("month")

		// ASSERT
		// Start should be first day of month at 00:00:00
		assert.Equal(t, 1, start.Day())
		assert.Equal(t, 0, start.Hour())
		assert.Equal(t, 0, start.Minute())
		assert.Equal(t, 0, start.Second())

		// Start should be current month
		now := time.Now().UTC()
		assert.Equal(t, now.Year(), start.Year())
		assert.Equal(t, now.Month(), start.Month())

		// End should be approximately now
		assert.WithinDuration(t, now, end, time.Second)

		// Start should be before or equal to end
		assert.True(t, start.Before(end) || start.Equal(end))
	})

	t.Run("unknown period", func(t *testing.T) {
		// ACT
		start, end := bot.getPeriodDates("unknown")

		// ASSERT
		// Start should be zero time
		assert.True(t, start.IsZero())

		// End should be approximately now
		now := time.Now().UTC()
		assert.WithinDuration(t, now, end, time.Second)
	})
}
