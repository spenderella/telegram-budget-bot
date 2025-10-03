package services

import (
	"database/sql"
	"log"
	"time"

	"telegram-finance-bot/internal/constants"
	"telegram-finance-bot/internal/errors"
	"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/repositories"
)

type ExpenseService struct {
	Repository      *repositories.ExpenseRepository
	UserService     *UserService
	CategoryService *CategoryService
}

func NewExpenseService(
	repository *repositories.ExpenseRepository, userService *UserService, categoryService *CategoryService) *ExpenseService {
	return &ExpenseService{
		Repository:      repository,
		UserService:     userService,
		CategoryService: categoryService,
	}
}

func (s *ExpenseService) AddExpense(telegramID int64, username string, amount float64, categoryName string, date time.Time) error {

	currency := constants.DefaultCurrency
	user, err := s.UserService.GetOrCreate(telegramID, username)
	if err != nil {
		return err
	}

	category, err := s.CategoryService.GetCategory(categoryName)
	if err == sql.ErrNoRows {
		return errors.ErrInvalidCategoryName
	}

	if err != nil {
		return err
	}

	expense := models.Expense{
		UserID:   user.ID,
		Amount:   amount,
		Category: category,
		Date:     date,
		Currency: currency,
	}

	err = s.Repository.Save(expense)
	if err != nil {
		log.Printf("Error saving expense for user %d: %v", telegramID, err)
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
