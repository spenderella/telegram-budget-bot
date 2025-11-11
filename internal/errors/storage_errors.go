package errors

import (
	"fmt"
)

func ErrFailedToGetData(err error) error {
	return fmt.Errorf("failed to get data: %w", err)
}

func ErrFailedToSaveExpenses(err error) error {
	return fmt.Errorf("failed to save expenses: %w", err)
}
