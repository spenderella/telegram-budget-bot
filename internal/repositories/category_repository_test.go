package repositories

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CategoryRepositoryTestSuite struct {
	RepositoryTestSuite
}

func (s *CategoryRepositoryTestSuite) Test_GetCategory_Success() {
	// ACT
	category, err := s.CategoryRepo.GetCategory("food")

	// ASSERT
	require.NoError(s.T(), err)
	assert.NotZero(s.T(), category.ID, "Category ID should be retrieved")
	assert.Equal(s.T(), "food", category.Name)
}

func (s *CategoryRepositoryTestSuite) Test_GetCategory_NotFound() {
	// ACT
	category, err := s.CategoryRepo.GetCategory("nonexistent")

	// ASSERT
	require.Error(s.T(), err)
	assert.Equal(s.T(), sql.ErrNoRows, err, "Should return sql.ErrNoRows for non-existing category")
	assert.Nil(s.T(), category, "Category should be nil when not found")
}

func (s *CategoryRepositoryTestSuite) Test_GetCategories_Success() {
	// ARRANGE
	expectedNames := []string{
		"food",
		"transportation",
		"house",
		"health",
		"entertainment",
		"personal",
		"other",
	}

	// ACT
	categories, err := s.CategoryRepo.GetCategories()

	var categoryNames []string
	for _, cat := range categories {
		categoryNames = append(categoryNames, cat.Name)
	}

	// ASSERT
	require.NoError(s.T(), err)
	assert.Len(s.T(), categories, 7, "Should have 7 default categories")
	assert.ElementsMatch(s.T(), expectedNames, categoryNames)
}

// TestCategoryRepositoryTestSuite runs the test suite
func TestCategoryRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryRepositoryTestSuite))
}
