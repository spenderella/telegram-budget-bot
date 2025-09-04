package main

import (
	"log"
	"os"

	"telegram-finance-bot/internal/api"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is not set")
	}

	bot, err := api.NewBudgetBot(token)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	bot.Start()
}
