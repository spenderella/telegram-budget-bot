package main

import (
	"log"

	"github.com/joho/godotenv"

	"telegram-finance-bot/internal/api"
	"telegram-finance-bot/internal/configs"
	"telegram-finance-bot/internal/repositories"
	"telegram-finance-bot/internal/services"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {

	config, err := configs.LoadConfig("internal/configs/configs.json")

	if err != nil {
		log.Fatal("Failed to read configs:", err)
	}

	expenseRepo := repositories.NewExpenseRepository()
	expenseService := services.NewExpenseService(expenseRepo)

	bot, err := api.NewBudgetBot(config, expenseService)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	bot.Start()
}
