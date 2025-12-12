package repositories

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// UserRepositoryTestSuite extends RepositoryTestSuite
type UserRepositoryTestSuite struct {
	RepositoryTestSuite
}

// TestCreateSuccess tests successful user creation
func (s *UserRepositoryTestSuite) TestCreateSuccess() {
	// ACT
	user, err := s.UserRepo.Create(12345, "testuser")

	// ASSERT
	require.NoError(s.T(), err)
	assert.NotZero(s.T(), user.ID, "User ID should be generated")
	assert.Equal(s.T(), int64(12345), user.TelegramID)
	assert.Equal(s.T(), "testuser", user.Username)
}

// TestCreateDuplicateTelegramID tests that duplicate telegram_id fails
func (s *UserRepositoryTestSuite) TestCreateDuplicateTelegramID() {
	// ARRANGE - create first user
	_, err := s.UserRepo.Create(99999, "user1")
	require.NoError(s.T(), err)

	// ACT - try to create user with same telegram_id
	_, err = s.UserRepo.Create(99999, "user2")

	// ASSERT
	require.Error(s.T(), err, "Should fail on duplicate telegram_id")
	assert.Contains(s.T(), err.Error(), "duplicate key value", "Should be unique constraint violation")
}

// TestGetUserExisting tests retrieving an existing user
func (s *UserRepositoryTestSuite) TestGetUserExisting() {
	// ARRANGE - create a user first
	createdUser, err := s.UserRepo.Create(54321, "existinguser")
	require.NoError(s.T(), err)

	// ACT
	user, err := s.UserRepo.GetUser(54321)

	// ASSERT
	require.NoError(s.T(), err)
	assert.Equal(s.T(), createdUser.ID, user.ID)
	assert.Equal(s.T(), int64(54321), user.TelegramID)
	assert.Equal(s.T(), "existinguser", user.Username)
}

// TestGetUserNotFound tests that non-existing user returns sql.ErrNoRows
func (s *UserRepositoryTestSuite) TestGetUserNotFound() {
	// ACT
	user, err := s.UserRepo.GetUser(99999)

	// ASSERT
	require.Error(s.T(), err)
	assert.Equal(s.T(), sql.ErrNoRows, err, "Should return sql.ErrNoRows for non-existing user")
	assert.Nil(s.T(), user, "User should be nil when not found")
}

// TestIntegrationCreateAndRetrieveMultiple verifies full flow with multiple users
func (s *UserRepositoryTestSuite) TestIntegrationCreateAndRetrieveMultiple() {
	// Create multiple users
	users := []struct {
		telegramID int64
		username   string
	}{
		{11111, "user1"},
		{22222, "user2"},
		{33333, "user3"},
	}

	createdIDs := make(map[int64]int)

	for _, u := range users {
		created, err := s.UserRepo.Create(u.telegramID, u.username)
		require.NoError(s.T(), err, "Failed to create user %s", u.username)
		createdIDs[u.telegramID] = created.ID
	}

	// Retrieve and verify each user
	for _, u := range users {
		retrieved, err := s.UserRepo.GetUser(u.telegramID)
		require.NoError(s.T(), err, "Failed to get user with telegram_id %d", u.telegramID)
		assert.Equal(s.T(), createdIDs[u.telegramID], retrieved.ID)
		assert.Equal(s.T(), u.telegramID, retrieved.TelegramID)
		assert.Equal(s.T(), u.username, retrieved.Username)
	}
}

// TestUserRepositoryTestSuite runs the test suite
func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
