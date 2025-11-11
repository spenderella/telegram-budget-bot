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

func ErrMissingEnvVars(varNames []string) error {
	return fmt.Errorf("missing required environment variables: %v", varNames)
}
