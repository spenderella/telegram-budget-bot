package errors

import (
	"errors"
)

var (
	ErrInvalidCategoryName = errors.New("category doesn't exist yet")
)
