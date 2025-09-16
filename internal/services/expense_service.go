package services

import (
	"fmt"
	"log"
	"time"

	"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/repositories"
)

type ExpenseService struct {
	Repository *repositories.ExpenseRepository
}

func (s *ExpenseService) AddExpense(userID int64, amount float64, category string, date time.Time) error {

	currency := "RUB" //TBD: currency from user's settings or message

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
		return fmt.Errorf("failed to save expense: %w", err)
	}
	return nil
}
