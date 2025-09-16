package repositories

import (
	"log"

	"telegram-finance-bot/internal/models"
)

type ExpenseRepository struct {
	Expenses map[int64][]models.Expense // userID -> список трат
}

func (r *ExpenseRepository) Save(expense models.Expense) error {

	if r.Expenses[expense.UserID] == nil {
		log.Printf("Creating new slice for user %d", expense.UserID)
		r.Expenses[expense.UserID] = []models.Expense{}
	}

	r.Expenses[expense.UserID] = append(r.Expenses[expense.UserID], expense)

	log.Printf("=== AFTER SAVE ===")
	log.Printf("Length after: %d", len(r.Expenses[expense.UserID]))
	log.Printf("All expenses: %v", r.Expenses[expense.UserID])

	return nil
}
