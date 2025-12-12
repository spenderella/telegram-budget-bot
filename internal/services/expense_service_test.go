package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	appErrors "telegram-finance-bot/internal/errors"
	"telegram-finance-bot/internal/models"
)

func TestExpenseServiceAddExpense(t *testing.T) {
	// Create a "database" of available categories
	availableCategories := generateTestCategories()

	// Helper to create a standard user service mock
	createMockUserService := func() *MockUserService {
		return &MockUserService{
			GetOrCreateFunc: func(telegramID int64, username string) (*models.User, error) {
				return &models.User{
					ID:         1,
					TelegramID: telegramID,
					Username:   username,
				}, nil
			},
		}
	}

	// Helper to create a category service mock with our test data
	createMockCategoryService := func() *MockCategoryService {
		return &MockCategoryService{
			GetCategoryFunc: func(name string) (*models.Category, error) {
				return findCategory(availableCategories, name)
			},
		}
	}

	t.Run("category not found", func(t *testing.T) {
		// ARRANGE
		mockUserService := createMockUserService()
		mockCategoryService := createMockCategoryService()
		mockRepository := &MockExpenseRepository{}

		service := NewExpenseService(mockRepository, mockUserService, mockCategoryService)

		// ACT
		err := service.AddExpense(12345, "testuser", 100.0, "books")

		// ASSERT
		require.Error(t, err)
		assert.Equal(t, appErrors.ErrInvalidCategoryName, err)
	})

	t.Run("category found - success", func(t *testing.T) {
		// ARRANGE
		mockUserService := createMockUserService()
		mockCategoryService := createMockCategoryService()

		// Track if Save was called
		saveCalled := false
		var savedExpense models.Expense

		mockRepository := &MockExpenseRepository{
			SaveFunc: func(expense models.Expense) error {
				saveCalled = true
				savedExpense = expense
				return nil
			},
		}

		service := NewExpenseService(mockRepository, mockUserService, mockCategoryService)

		// ACT - try to add expense with existing category
		err := service.AddExpense(12345, "testuser", 100.0, "food")

		// ASSERT
		require.NoError(t, err)
		assert.True(t, saveCalled, "Repository.Save should have been called")
		assert.Equal(t, 100.0, savedExpense.Amount)
		assert.Equal(t, "food", savedExpense.Category.Name)
		assert.Equal(t, 1, savedExpense.UserID)
	})

	t.Run("database error", func(t *testing.T) {
		// ARRANGE
		mockUserService := createMockUserService()
		mockCategoryService := createMockCategoryService()

		saveCalled := false
		dbError := errors.New("database connection lost")

		mockRepository := &MockExpenseRepository{
			SaveFunc: func(expense models.Expense) error {
				saveCalled = true
				return dbError
			},
		}

		service := NewExpenseService(mockRepository, mockUserService, mockCategoryService)

		// ACT
		err := service.AddExpense(12345, "testuser", 100.0, "food")

		// ASSERT
		require.Error(t, err)
		assert.True(t, saveCalled, "Repository.Save should have been called")
		assert.Contains(t, err.Error(), "failed to save expenses")
		assert.Contains(t, err.Error(), "database connection lost")
	})

	t.Run("user service error", func(t *testing.T) {
		// ARRANGE
		userError := errors.New("user database unavailable")
		mockUserService := &MockUserService{
			GetOrCreateFunc: func(telegramID int64, username string) (*models.User, error) {
				return nil, userError
			},
		}
		mockCategoryService := createMockCategoryService()
		mockRepository := &MockExpenseRepository{}

		service := NewExpenseService(mockRepository, mockUserService, mockCategoryService)

		// ACT
		err := service.AddExpense(12345, "testuser", 100.0, "food")

		// ASSERT
		require.Error(t, err)
		assert.Equal(t, userError, err)
	})
}
func TestExpenseServiceGetExpenses(t *testing.T) {

	availableCategories := generateTestCategories()

	t.Run("get expenses success", func(t *testing.T) {
		// ARRANGE
		testExpenses := generateTestExpenses(5, 1, availableCategories)
		mockRepository := &MockExpenseRepository{
			GetExpensesFunc: func(filter models.ExpenseFilter) ([]models.Expense, error) {
				return testExpenses, nil
			},
		}

		service := NewExpenseService(mockRepository, nil, nil)

		// ACT

		filter := models.ExpenseFilter{}
		resultExpenses, err := service.GetExpenses(filter)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(t, testExpenses, resultExpenses)
	})

	t.Run("get expenses error", func(t *testing.T) {
		// ARRANGE
		repositoryError := errors.New("repository error")
		mockRepository := &MockExpenseRepository{
			GetExpensesFunc: func(filter models.ExpenseFilter) ([]models.Expense, error) {
				return nil, repositoryError
			},
		}

		service := NewExpenseService(mockRepository, nil, nil)

		// ACT
		_, err := service.GetExpenses(models.ExpenseFilter{})

		// ASSERT
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get data")
		assert.Contains(t, err.Error(), "repository error")
	})
}
func TestExpenseServiceGetStats(t *testing.T) {
	t.Run("get stats error", func(t *testing.T) {
		// ARRANGE
		repositoryError := errors.New("repository error")
		mockRepository := &MockExpenseRepository{
			GetStatFunc: func(filter models.ExpenseFilter) ([]models.CategoryExpenses, error) {
				return nil, repositoryError
			},
		}

		service := NewExpenseService(mockRepository, nil, nil)

		// ACT
		_, err := service.GetStat(models.ExpenseFilter{})

		// ASSERT
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get data")
		assert.Contains(t, err.Error(), "repository error")
	})

}
