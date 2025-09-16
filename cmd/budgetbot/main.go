package main

import (
	"log"

	"telegram-finance-bot/internal/api"
	"telegram-finance-bot/internal/configs"
	"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/repositories"
	"telegram-finance-bot/internal/services"
)

func main() {
	config, err := configs.LoadConfig("internal/configs/configs.json")

	if err != nil {
		log.Fatal("Failed to read configs:", err)
	}

	ER := repositories.ExpenseRepository{
		Expenses: make(map[int64][]models.Expense),
	}
	ES := services.ExpenseService{
		Repository: &ER,
	}

	bot, err := api.NewBudgetBot(config, &ES)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	bot.Start()
}
