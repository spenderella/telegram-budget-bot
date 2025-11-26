package services

import (
	"database/sql"
	"time"

	"telegram-finance-bot/internal/models"
)

// Mock for IUserService
type MockUserService struct {
	GetOrCreateFunc func(telegramID int64, username string) (*models.User, error)
}

func (m *MockUserService) GetOrCreate(telegramID int64, username string) (*models.User, error) {
	if m.GetOrCreateFunc != nil {
		return m.GetOrCreateFunc(telegramID, username)
	}
	return nil, nil
}

// Mock for ICategoryService
type MockCategoryService struct {
	GetCategoryFunc   func(name string) (*models.Category, error)
	GetCategoriesFunc func() ([]models.Category, error)
}

func (m *MockCategoryService) GetCategory(name string) (*models.Category, error) {
	if m.GetCategoryFunc != nil {
		return m.GetCategoryFunc(name)
	}
	return nil, nil
}

func (m *MockCategoryService) GetCategories() ([]models.Category, error) {
	if m.GetCategoriesFunc != nil {
		return m.GetCategoriesFunc()
	}
	return nil, nil
}

// Mock for IExpenseRepository
type MockExpenseRepository struct {
	SaveFunc        func(expense models.Expense) error
	GetExpensesFunc func(filter models.ExpenseFilter) ([]models.Expense, error)
	GetStatFunc     func(filter models.ExpenseFilter) ([]models.CategoryExpenses, error)
}

func (m *MockExpenseRepository) Save(expense models.Expense) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(expense)
	}
	return nil
}

func (m *MockExpenseRepository) GetExpenses(filter models.ExpenseFilter) ([]models.Expense, error) {
	if m.GetExpensesFunc != nil {
		return m.GetExpensesFunc(filter)
	}
	return nil, nil
}

func (m *MockExpenseRepository) GetStat(filter models.ExpenseFilter) ([]models.CategoryExpenses, error) {
	if m.GetStatFunc != nil {
		return m.GetStatFunc(filter)
	}
	return nil, nil
}
func generateTestCategories() []models.Category {
	categories := []models.Category{
		{ID: 1, Name: "food"},
		{ID: 2, Name: "transport"},
		{ID: 3, Name: "entertainment"},
	}
	return categories
}

// Helper function to find category by name (simulates DB search)
func findCategory(categories []models.Category, name string) (*models.Category, error) {
	for _, category := range categories {
		if category.Name == name {
			return &category, nil
		}
	}
	return nil, sql.ErrNoRows
}

// Helper function to generate test expenses
func generateTestExpenses(count int, userID int, categories []models.Category) []models.Expense {
	expenses := make([]models.Expense, count)
	now := time.Now()

	for i := 0; i < count; i++ {
		expenses[i] = models.Expense{
			Date:     now.Add(-time.Duration(i*24) * time.Hour), // each expense 1 day apart
			Amount:   float64((i + 1) * 5),                      // 5, 10, 15, etc.
			UserID:   userID,
			Currency: "EUR",
			Category: &categories[i%len(categories)], // rotate through categories
		}
	}

	return expenses
}
