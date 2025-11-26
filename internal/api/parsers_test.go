package api

import (
	"testing"

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
