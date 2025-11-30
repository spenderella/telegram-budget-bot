package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"telegram-finance-bot/internal/database"
)

// RepositoryTestSuite is a base test suite for all repository integration tests
type RepositoryTestSuite struct {
	suite.Suite
	DB               *sql.DB
	UserRepo         *UserRepository
	CategoryRepo     *CategoryRepository
	ExpenseRepo      *ExpenseRepository
}

// SetupSuite runs once before all tests - creates DB connection and runs migrations
func (s *RepositoryTestSuite) SetupSuite() {
	// Connection string for test database
	connStr := "host=localhost port=5433 user=testuser password=testpass dbname=telegram_bot_test sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	require.NoError(s.T(), err, "Failed to connect to test database")

	// Verify connection
	err = db.Ping()
	require.NoError(s.T(), err, "Failed to ping test database")

	// Run migrations
	err = database.RunMigrations(db)
	require.NoError(s.T(), err, "Failed to run migrations")

	s.DB = db

	// Initialize all repositories
	s.UserRepo = NewUserRepository(db)
	s.CategoryRepo = NewCategoryRepository(db)
	s.ExpenseRepo = NewExpenseRepository(db)
}

// TearDownSuite runs once after all tests - closes DB connection
func (s *RepositoryTestSuite) TearDownSuite() {
	if s.DB != nil {
		err := s.DB.Close()
		require.NoError(s.T(), err, "Failed to close database connection")
	}
}

// AfterTest runs after each test - cleans up all data
func (s *RepositoryTestSuite) AfterTest(suiteName, testName string) {
	// Delete in correct order due to foreign keys
	_, err := s.DB.Exec("DELETE FROM expenses")
	require.NoError(s.T(), err, "Failed to cleanup expenses")

	_, err = s.DB.Exec("DELETE FROM users")
	require.NoError(s.T(), err, "Failed to cleanup users")

	// Categories are seeded by migrations, we keep them
}
