package repositories

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"fmt"
	"os"
	"telegram-finance-bot/internal/database"
	"telegram-finance-bot/internal/errors"

	"github.com/joho/godotenv"
)

// RepositoryTestSuite is a base test suite for all repository integration tests
type RepositoryTestSuite struct {
	suite.Suite
	DB           *sql.DB
	UserRepo     *UserRepository
	CategoryRepo *CategoryRepository
	ExpenseRepo  *ExpenseRepository
}

// SetupSuite runs once before all tests - creates DB connection and runs migrations
func (s *RepositoryTestSuite) SetupSuite() {
	err := godotenv.Load("../../.env")
	require.NoError(s.T(), err, "Error loading .env file")

	// Connection string for test database
	requiredEnvVars := []string{"TEST_DB_HOST", "TEST_DB_PORT", "TEST_DB_USER", "TEST_DB_PASSWORD", "TEST_DB_NAME", "TEST_DB_SSLMODE"}

	envVars := make(map[string]string)
	var missingVars []string
	for _, varName := range requiredEnvVars {
		if os.Getenv(varName) == "" {
			missingVars = append(missingVars, varName)
		} else {
			envVars[varName] = os.Getenv(varName)
		}
	}

	if len(missingVars) > 0 {
		err := errors.ErrMissingEnvVars(missingVars)
		require.NoError(s.T(), err, "Failed to get required credentials to test database")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envVars["TEST_DB_HOST"],
		envVars["TEST_DB_PORT"],
		envVars["TEST_DB_USER"],
		envVars["TEST_DB_PASSWORD"],
		envVars["TEST_DB_NAME"],
		envVars["TEST_DB_SSLMODE"])

	db, err := sql.Open("postgres", dsn)
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
