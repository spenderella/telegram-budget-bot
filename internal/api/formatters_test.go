package api

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"telegram-finance-bot/internal/models"
)

func Test_formatExpenses(t *testing.T) {
	bot := &BudgetBot{}

	t.Run("empty expenses", func(t *testing.T) {
		// ARRANGE
		expenses := []models.Expense{}

		// ACT
		result := bot.formatExpenses(expenses)

		// ASSERT
		assert.Equal(t, "📊 Expenses not found", result)
	})

	t.Run("single expense", func(t *testing.T) {
		// ARRANGE
		date := time.Date(2025, 11, 27, 14, 30, 0, 0, time.UTC)
		expenses := []models.Expense{
			{
				Amount:   100.50,
				Currency: "EUR",
				Category: &models.Category{ID: 1, Name: "food"},
				Date:     date,
			},
		}

		// ACT
		result := bot.formatExpenses(expenses)

		// ASSERT
		assert.Contains(t, result, "📊 Found 1 expenses:")
		assert.Contains(t, result, "Amount: 100.50 EUR")
		assert.Contains(t, result, "Category: food")
		assert.Contains(t, result, "27.11.2025 14:30")
		assert.Contains(t, result, "Total amount: 100.50")
	})

	t.Run("multiple expenses", func(t *testing.T) {
		// ARRANGE
		date1 := time.Date(2025, 11, 27, 14, 30, 0, 0, time.UTC)
		date2 := time.Date(2025, 11, 28, 10, 15, 0, 0, time.UTC)
		expenses := []models.Expense{
			{
				Amount:   100.50,
				Currency: "EUR",
				Category: &models.Category{ID: 1, Name: "food"},
				Date:     date1,
			},
			{
				Amount:   50.25,
				Currency: "EUR",
				Category: &models.Category{ID: 2, Name: "transport"},
				Date:     date2,
			},
		}

		// ACT
		result := bot.formatExpenses(expenses)

		// ASSERT
		assert.Contains(t, result, "📊 Found 2 expenses:")
		assert.Contains(t, result, "Amount: 100.50 EUR")
		assert.Contains(t, result, "Category: food")
		assert.Contains(t, result, "Amount: 50.25 EUR")
		assert.Contains(t, result, "Category: transport")
		assert.Contains(t, result, "Total amount: 150.75")

		// Verify structure: each expense should have Amount, Category, and Date
		expenseBlocks := strings.Split(result, "\n\n")
		// Should have: header + 2 expenses + total = at least 3 blocks
		assert.GreaterOrEqual(t, len(expenseBlocks), 3)
	})

	t.Run("total calculation", func(t *testing.T) {
		// ARRANGE
		date := time.Date(2025, 11, 27, 14, 30, 0, 0, time.UTC)
		expenses := []models.Expense{
			{Amount: 10.10, Currency: "EUR", Category: &models.Category{Name: "a"}, Date: date},
			{Amount: 20.20, Currency: "EUR", Category: &models.Category{Name: "b"}, Date: date},
			{Amount: 30.30, Currency: "EUR", Category: &models.Category{Name: "c"}, Date: date},
		}

		// ACT
		result := bot.formatExpenses(expenses)

		// ASSERT
		// 10.10 + 20.20 + 30.30 = 60.60
		assert.Contains(t, result, "Total amount: 60.60")
	})
}

func Test_formatCategoriesStat(t *testing.T) {
	bot := &BudgetBot{}

	t.Run("empty statistics", func(t *testing.T) {
		// ARRANGE
		stat := []models.CategoryExpenses{}

		// ACT
		result := bot.formatCategoriesStat(stat)

		// ASSERT
		assert.Equal(t, "📊 Data not found", result)
	})

	t.Run("single category", func(t *testing.T) {
		// ARRANGE
		stat := []models.CategoryExpenses{
			{
				Category:      &models.Category{ID: 1, Name: "food"},
				Total:         250.75,
				TotalCurrency: "EUR",
			},
		}

		// ACT
		result := bot.formatCategoriesStat(stat)

		// ASSERT
		assert.Contains(t, result, "📊 Statistics for categories:")
		assert.Contains(t, result, "Category: food")
		assert.Contains(t, result, "Total: 250.75 EUR")
	})

	t.Run("multiple categories", func(t *testing.T) {
		// ARRANGE
		stat := []models.CategoryExpenses{
			{
				Category:      &models.Category{ID: 1, Name: "food"},
				Total:         250.75,
				TotalCurrency: "EUR",
			},
			{
				Category:      &models.Category{ID: 2, Name: "transport"},
				Total:         100.50,
				TotalCurrency: "EUR",
			},
			{
				Category:      &models.Category{ID: 3, Name: "entertainment"},
				Total:         75.25,
				TotalCurrency: "EUR",
			},
		}

		// ACT
		result := bot.formatCategoriesStat(stat)

		// ASSERT
		assert.Contains(t, result, "📊 Statistics for categories:")
		assert.Contains(t, result, "Category: food")
		assert.Contains(t, result, "Total: 250.75 EUR")
		assert.Contains(t, result, "Category: transport")
		assert.Contains(t, result, "Total: 100.50 EUR")
		assert.Contains(t, result, "Category: entertainment")
		assert.Contains(t, result, "Total: 75.25 EUR")

		// Verify all three categories are present
		lines := strings.Split(result, "\n")
		categoryLines := 0
		for _, line := range lines {
			if strings.Contains(line, "Category:") {
				categoryLines++
			}
		}
		assert.Equal(t, 3, categoryLines, "Should have 3 category lines")
	})

	t.Run("formatting precision", func(t *testing.T) {
		// ARRANGE
		stat := []models.CategoryExpenses{
			{
				Category:      &models.Category{ID: 1, Name: "test"},
				Total:         123.456, // Should be formatted to 2 decimal places
				TotalCurrency: "EUR",
			},
		}

		// ACT
		result := bot.formatCategoriesStat(stat)

		// ASSERT
		assert.Contains(t, result, "Total: 123.46 EUR", "Should format to 2 decimal places")
	})
}
