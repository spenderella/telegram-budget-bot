package services

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"telegram-finance-bot/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCategoryServiceGetCategory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// ARRANGE
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedCategory := &models.Category{ID: 1, Name: "food"}

		mockRepo := NewMockICategoryRepository(ctrl)

		mockRepo.EXPECT().
			GetCategory("food").
			Return(expectedCategory, nil).
			Times(1)

		service := NewCategoryService(mockRepo)

		// ACT
		category, err := service.GetCategory("food")

		// ASSERT
		require.NoError(t, err)
		assert.Equal(t, expectedCategory, category)

	})

	t.Run("error", func(t *testing.T) {
		// ARRANGE
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedError := errors.New("repository error")

		mockRepo := NewMockICategoryRepository(ctrl)

		mockRepo.EXPECT().
			GetCategory("food").
			Return(nil, expectedError).
			Times(1)

		service := NewCategoryService(mockRepo)

		// ACT
		_, err := service.GetCategory("food")

		// ASSERT
		require.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestCategoryServiceGetCategories(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// ARRANGE
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedCategories := generateTestCategories()

		mockRepo := NewMockICategoryRepository(ctrl)

		mockRepo.EXPECT().
			GetCategories().
			Return(expectedCategories, nil).
			Times(1)

		service := NewCategoryService(mockRepo)

		// ACT
		categories, err := service.GetCategories()

		// ASSERT
		require.NoError(t, err)
		assert.Equal(t, expectedCategories, categories)

	})

	t.Run("error", func(t *testing.T) {
		// ARRANGE
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expectedError := errors.New("repository error")

		mockRepo := NewMockICategoryRepository(ctrl)

		mockRepo.EXPECT().
			GetCategories().
			Return(nil, expectedError).
			Times(1)

		service := NewCategoryService(mockRepo)

		// ACT
		_, err := service.GetCategories()

		// ASSERT

		require.Error(t, err)
		assert.Equal(t, expectedError, err)

	})
}
