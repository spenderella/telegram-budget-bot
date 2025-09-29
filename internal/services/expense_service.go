package services

import (
	"log"
	"time"

	"telegram-finance-bot/internal/errors"
	"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/repositories"
)

type ExpenseService struct {
	Repository *repositories.ExpenseRepository
}

func NewExpenseService(repository *repositories.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		Repository: repository,
	}
}

func (s *ExpenseService) AddExpense(userID int64, amount float64, category string, date time.Time) error {

	currency := defaultCurrency
	expense := models.Expense{
		UserID:   userID,
		Amount:   amount,
		Category: category,
		Date:     date,
		Currency: currency,
	}

	err := s.Repository.Save(expense)
	if err != nil {
		log.Printf("Error saving expense for user %d: %v", userID, err)
		return errors.ErrFailedToSaveExpenses(err)
	}
	return nil
}

func (s *ExpenseService) GetExpenses(filter models.ExpenseFilter) ([]models.Expense, error) {
	expenses, err := s.Repository.GetExpenses(filter)
	if err != nil {
		return nil, errors.ErrFailedToGetExpenses(err)
	}

	return expenses, nil
}
