package errors

import (
	"fmt"
)

func ErrFailedToGetExpenses(err error) error {
	return fmt.Errorf("failed to get expenses: %w", err)
}

func ErrFailedToSaveExpenses(err error) error {
	return fmt.Errorf("failed to save expenses: %w", err)
}
