package api

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BudgetBot struct {
	token string
	debug bool
	api   *tgbotapi.BotAPI
}

// Конструктор
func NewBudgetBot(token string) (*BudgetBot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &BudgetBot{
		token: token,
		debug: true,
		api:   botAPI,
	}, nil
}

func (bot *BudgetBot) Start() {
	bot.api.Debug = bot.debug
	log.Printf("Authorized on account %s", bot.api.Self.UserName)

	// Создаем конфигурацию для получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получаем канал для обновлений
	updates := bot.api.GetUpdatesChan(u)

	log.Println("Bot is running. Send /start to test...")
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			bot.handleMessage(&update)
		}
	}
}
