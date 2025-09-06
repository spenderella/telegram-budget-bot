package services

import "telegram-finance-bot/internal/repositories"

type ExpenseService struct {
	repository *repositories.ExpenseRepository
}
