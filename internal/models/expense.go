package models

import "time"

type Expense struct {
	Date     time.Time
	Amount   float64
	UserID   int64
	Currency string
	Category string
}

type ExpenseFilter struct {
	UserID int64
	//Category *string
	//DateFrom *time.Time
	//DateTo   *time.Time
	Limit *int
}
