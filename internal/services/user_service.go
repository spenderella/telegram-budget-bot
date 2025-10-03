package services

import (
	"database/sql"

	"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetOrCreate(telegramID int64, username string) (*models.User, error) {
	user, err := s.userRepo.GetUser(telegramID)

	if err == sql.ErrNoRows {
		return s.userRepo.Create(telegramID, username)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
