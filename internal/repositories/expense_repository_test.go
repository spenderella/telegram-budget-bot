package repositories

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"telegram-finance-bot/internal/models"
)

type ExpenseRepositoryTestSuite struct {
	RepositoryTestSuite
	testUser1     *models.User
	testUser2     *models.User
	foodCategory  *models.Category
	houseCategory *models.Category
}

func (s *ExpenseRepositoryTestSuite) SetupSuite() {
	// Call parent SetupSuite to initialize DB and repos
	s.RepositoryTestSuite.SetupSuite()

	// Create test users
	var err error
	s.testUser1, err = s.UserRepo.Create(12345, "testuser1")
	require.NoError(s.T(), err)

	s.testUser2, err = s.UserRepo.Create(67890, "testuser2")
	require.NoError(s.T(), err)

	// Get test categories
	s.foodCategory, err = s.CategoryRepo.GetCategory("food")
	require.NoError(s.T(), err)

	s.houseCategory, err = s.CategoryRepo.GetCategory("house")
	require.NoError(s.T(), err)

	// Create some test expenses
	expenses := []models.Expense{
		{
			UserID:   s.testUser1.ID,
			Category: s.foodCategory,
			Amount:   50.00,
			Currency: "USD",
		},
		{
			UserID:   s.testUser1.ID,
			Category: s.foodCategory,
			Amount:   30.00,
			Currency: "USD",
		},
		{
			UserID:   s.testUser1.ID,
			Category: s.houseCategory,
			Amount:   200.00,
			Currency: "USD",
		},
		{
			UserID:   s.testUser2.ID,
			Category: s.foodCategory,
			Amount:   100.00,
			Currency: "USD",
		},
	}

	for _, exp := range expenses {
		err := s.ExpenseRepo.Save(exp)
		require.NoError(s.T(), err)
	}
}

// AfterTest overrides parent AfterTest to NOT clean up data
// since we use fixtures created in SetupSuite
func (s *ExpenseRepositoryTestSuite) AfterTest(suiteName, testName string) {
	// Do nothing - keep test data for all tests
}

func (s *ExpenseRepositoryTestSuite) TestSaveInvalidUserID() {
	// ARRANGE
	expense := models.Expense{
		UserID:   99999, // non-existent user
		Category: s.foodCategory,
		Amount:   100.50,
		Currency: "USD",
	}

	// ACT
	err := s.ExpenseRepo.Save(expense)

	// ASSERT
	require.Error(s.T(), err, "Should fail with foreign key constraint violation")
}

func (s *ExpenseRepositoryTestSuite) TestGetExpensesUser1() {
	// ARRANGE
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	tomorrow := time.Now().UTC().Add(24 * time.Hour)
	limit := 10

	filter := models.ExpenseFilter{
		UserTgID: s.testUser1.TelegramID,
		DateFrom: &yesterday,
		DateTo:   &tomorrow,
		Limit:    &limit,
	}

	// ACT
	result, err := s.ExpenseRepo.GetExpenses(filter)

	// ASSERT
	require.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(result), 3, "User1 should have at least 3 expenses from fixtures")

	// Check that all expenses belong to user1
	for _, exp := range result {
		assert.Equal(s.T(), s.testUser1.ID, exp.UserID)
	}
}

func (s *ExpenseRepositoryTestSuite) TestGetExpensesUser2() {
	// ARRANGE
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	tomorrow := time.Now().UTC().Add(24 * time.Hour)
	limit := 10

	filter := models.ExpenseFilter{
		UserTgID: s.testUser2.TelegramID,
		DateFrom: &yesterday,
		DateTo:   &tomorrow,
		Limit:    &limit,
	}

	// ACT
	result, err := s.ExpenseRepo.GetExpenses(filter)

	// ASSERT
	require.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(result), 1, "User2 should have at least 1 expense from fixtures")
	assert.Equal(s.T(), s.testUser2.ID, result[0].UserID)
}

func (s *ExpenseRepositoryTestSuite) TestGetExpensesEmptyResult() {
	// ARRANGE
	// Query far in the past
	past := time.Now().UTC().Add(-365 * 24 * time.Hour)
	longAgo := time.Now().UTC().Add(-366 * 24 * time.Hour)
	limit := 10

	filter := models.ExpenseFilter{
		UserTgID: s.testUser1.TelegramID,
		DateFrom: &longAgo,
		DateTo:   &past,
		Limit:    &limit,
	}

	// ACT
	result, err := s.ExpenseRepo.GetExpenses(filter)

	// ASSERT
	require.NoError(s.T(), err)
	assert.Empty(s.T(), result, "Should return empty slice when no expenses in date range")
}

func (s *ExpenseRepositoryTestSuite) TestGetExpensesRespectLimit() {
	// ARRANGE
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	tomorrow := time.Now().UTC().Add(24 * time.Hour)
	limit := 2

	filter := models.ExpenseFilter{
		UserTgID: s.testUser1.TelegramID,
		DateFrom: &yesterday,
		DateTo:   &tomorrow,
		Limit:    &limit,
	}

	// ACT
	result, err := s.ExpenseRepo.GetExpenses(filter)

	// ASSERT
	require.NoError(s.T(), err)
	assert.LessOrEqual(s.T(), len(result), 2, "Should respect limit and return at most 2 expenses")
}

func (s *ExpenseRepositoryTestSuite) TestGetExpensesOrderedByDate() {
	// ARRANGE
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	tomorrow := time.Now().UTC().Add(24 * time.Hour)
	limit := 10

	filter := models.ExpenseFilter{
		UserTgID: s.testUser1.TelegramID,
		DateFrom: &yesterday,
		DateTo:   &tomorrow,
		Limit:    &limit,
	}

	// ACT
	result, err := s.ExpenseRepo.GetExpenses(filter)

	// ASSERT
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result)

	// Check that results are ordered by date DESC
	for i := 0; i < len(result)-1; i++ {
		assert.True(s.T(),
			result[i].Date.After(result[i+1].Date) || result[i].Date.Equal(result[i+1].Date),
			"Expenses should be ordered by date DESC")
	}
}

func (s *ExpenseRepositoryTestSuite) TestGetStatUser1() {
	// ARRANGE
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	tomorrow := time.Now().UTC().Add(24 * time.Hour)
	limit := 10

	filter := models.ExpenseFilter{
		UserTgID: s.testUser1.TelegramID,
		DateFrom: &yesterday,
		DateTo:   &tomorrow,
		Limit:    &limit,
	}

	// ACT
	stats, err := s.ExpenseRepo.GetStat(filter)

	// ASSERT
	require.NoError(s.T(), err)
	assert.NotEmpty(s.T(), stats, "Should have statistics for user1")

	// Find food and house stats
	var foodStat, houseStat *models.CategoryExpenses
	for i := range stats {
		if stats[i].Category.Name == "food" {
			foodStat = &stats[i]
		} else if stats[i].Category.Name == "house" {
			houseStat = &stats[i]
		}
	}

	require.NotNil(s.T(), foodStat, "Should have food category stats")
	require.NotNil(s.T(), houseStat, "Should have house category stats")

	assert.GreaterOrEqual(s.T(), foodStat.Total, 80.00, "Food total should be at least 50 + 30")
	assert.Equal(s.T(), "USD", foodStat.TotalCurrency)
	assert.GreaterOrEqual(s.T(), houseStat.Total, 200.00, "House total should be at least 200")
}

func (s *ExpenseRepositoryTestSuite) TestGetStatEmptyResult() {
	// ARRANGE
	// Query far in the past
	past := time.Now().UTC().Add(-365 * 24 * time.Hour)
	longAgo := time.Now().UTC().Add(-366 * 24 * time.Hour)
	limit := 10

	filter := models.ExpenseFilter{
		UserTgID: s.testUser1.TelegramID,
		DateFrom: &longAgo,
		DateTo:   &past,
		Limit:    &limit,
	}

	// ACT
	stats, err := s.ExpenseRepo.GetStat(filter)

	// ASSERT
	require.NoError(s.T(), err)
	assert.Empty(s.T(), stats, "Should return empty slice when no expenses in date range")
}

func (s *ExpenseRepositoryTestSuite) TestGetStatIsolatedByUser() {
	// ARRANGE
	yesterday := time.Now().UTC().Add(-24 * time.Hour)
	tomorrow := time.Now().UTC().Add(24 * time.Hour)
	limit := 10

	filter1 := models.ExpenseFilter{
		UserTgID: s.testUser1.TelegramID,
		DateFrom: &yesterday,
		DateTo:   &tomorrow,
		Limit:    &limit,
	}

	filter2 := models.ExpenseFilter{
		UserTgID: s.testUser2.TelegramID,
		DateFrom: &yesterday,
		DateTo:   &tomorrow,
		Limit:    &limit,
	}

	// ACT
	stats1, err := s.ExpenseRepo.GetStat(filter1)
	require.NoError(s.T(), err)

	stats2, err := s.ExpenseRepo.GetStat(filter2)
	require.NoError(s.T(), err)

	// ASSERT
	// User1 should have multiple categories
	assert.GreaterOrEqual(s.T(), len(stats1), 2, "User1 should have stats for at least 2 categories")

	// User2 should have only food category
	assert.Equal(s.T(), 1, len(stats2), "User2 should have stats for 1 category")
	assert.Equal(s.T(), "food", stats2[0].Category.Name)
	assert.GreaterOrEqual(s.T(), stats2[0].Total, 100.00)
}

// TestExpenseRepositoryTestSuite runs the test suite
func TestExpenseRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ExpenseRepositoryTestSuite))
}
