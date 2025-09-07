package models

import "time"

type Expense struct {
	Date     time.Time
	Amount   float64
	UserID   int
	Currency string
	Category string
}
