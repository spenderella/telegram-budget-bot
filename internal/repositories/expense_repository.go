package repositories

import (
	"database/sql"
	"log"

	"telegram-finance-bot/internal/models"
)

type ExpenseRepository struct {
	Expenses map[int64][]models.Expense //deprecate
	db       *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{
		Expenses: make(map[int64][]models.Expense),
		db:       db,
	}
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

func (r *ExpenseRepository) GetExpenses(filter models.ExpenseFilter) ([]models.Expense, error) {
	userExpenses, exists := r.Expenses[filter.UserID]

	if !exists {
		return []models.Expense{}, nil
	}

	var result []models.Expense

	for _, expense := range userExpenses {

		result = append(result, expense)

		if filter.Limit == nil || len(result) >= *filter.Limit {
			break
		}
	}
	return result, nil
}
