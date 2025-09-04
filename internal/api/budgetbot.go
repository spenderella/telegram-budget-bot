package api

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Создаем конфигурацию для получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получаем канал для обновлений
	updates := bot.GetUpdatesChan(u)

	log.Println("Bot is running. Send /start to test...")
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Text {
			case "/start":
				handleStart(bot, &update)
			case "/help":
				handleHelp(bot, &update)
			default:
				handleUnknown(bot, &update)
			}
		}
	}
}
