package repositories

import (
	"log"

	"telegram-finance-bot/internal/models"
)

type ExpenseRepository struct {
	expenses map[int64][]models.Expense // userID -> список трат
}

func NewExpenseRepository() *ExpenseRepository {
	return &ExpenseRepository{
		expenses: make(map[int64][]models.Expense),
	}
}

func (r *ExpenseRepository) Save(expense models.Expense) error {

	if r.expenses[expense.UserID] == nil {
		log.Printf("Creating new slice for user %d", expense.UserID)
		r.expenses[expense.UserID] = []models.Expense{}
	}

	r.expenses[expense.UserID] = append(r.expenses[expense.UserID], expense)

	log.Printf("=== AFTER SAVE ===")
	log.Printf("Length after: %d", len(r.expenses[expense.UserID]))
	log.Printf("All expenses: %v", r.expenses[expense.UserID])

	return nil
}
