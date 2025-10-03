package models

import "time"

type Expense struct {
	Date     time.Time
	Amount   float64
	UserID   int
	Currency string
	Category *Category
}

type ExpenseFilter struct {
	UserID int64
	//Category *string
	//DateFrom *time.Time
	//DateTo   *time.Time
	Limit *int
}
