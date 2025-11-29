package services

import (
	"errors"
	"testing"

	"database/sql"

	"go.uber.org/mock/gomock"

	"telegram-finance-bot/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testTelegramID int64 = 1234

func TestUserService_GetOrCreate(t *testing.T) {
	t.Run("user exist", func(t *testing.T) {
		// ARRANGE
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedUser := &models.User{ID: 1, TelegramID: testTelegramID, Username: "username"}

		mockRepo := NewMockIUserRepository(ctrl)

		mockRepo.EXPECT().
			GetUser(testTelegramID).
			Return(expectedUser, nil).
			Times(1)

		mockRepo.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)

		service := NewUserService(mockRepo)

		// ACT
		user, err := service.GetOrCreate(testTelegramID, "username")

		// ASSERT
		require.NoError(t, err)
		assert.Equal(t, expectedUser, user)

	})

	t.Run("repository error", func(t *testing.T) {
		// ARRANGE
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedError := errors.New("repository error")

		mockRepo := NewMockIUserRepository(ctrl)

		mockRepo.EXPECT().
			GetUser(testTelegramID).
			Return(nil, expectedError).
			Times(1)

		mockRepo.EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)

		service := NewUserService(mockRepo)

		// ACT
		_, err := service.GetOrCreate(testTelegramID, "username")

		// ASSERT
		require.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
	t.Run("new user created", func(t *testing.T) {
		// ARRANGE
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedUser := &models.User{ID: 1, TelegramID: testTelegramID, Username: "username"}
		expectedError := sql.ErrNoRows

		mockRepo := NewMockIUserRepository(ctrl)

		mockRepo.EXPECT().
			GetUser(testTelegramID).
			Return(nil, expectedError).
			Times(1)

		mockRepo.EXPECT().
			Create(testTelegramID, "username").
			Return(expectedUser, nil).
			Times(1)

		service := NewUserService(mockRepo)

		// ACT
		user, err := service.GetOrCreate(testTelegramID, "username")

		// ASSERT
		require.NoError(t, err)
		assert.Equal(t, expectedUser, user)

	})

	t.Run("repository error", func(t *testing.T) {
		// ARRANGE
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedError := errors.New("repository error")

		mockRepo := NewMockIUserRepository(ctrl)

		mockRepo.EXPECT().
			GetUser(testTelegramID).
			Return(nil, sql.ErrNoRows).
			Times(1)

		mockRepo.EXPECT().
			Create(testTelegramID, "username").
			Return(nil, expectedError).
			Times(1)

		service := NewUserService(mockRepo)

		// ACT
		_, err := service.GetOrCreate(testTelegramID, "username")

		// ASSERT
		require.Error(t, err)
		assert.Equal(t, expectedError, err)

	})
}
