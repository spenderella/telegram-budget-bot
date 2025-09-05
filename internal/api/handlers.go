package api

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// / Обработка входящих сообщений
func (bot *BudgetBot) handleMessage(update *tgbotapi.Update) {
	command := strings.ToLower(update.Message.Text)

	switch command {
	case "/start":
		bot.commandStart(update)
	case "/help":
		bot.commandHelp(update)
	default:
		bot.commandUnknown(update)
	}
}

// Команда /start
func (bot *BudgetBot) commandStart(update *tgbotapi.Update) {
	username := update.Message.From.UserName
	if username == "" {
		username = update.Message.From.FirstName
	}

	text := fmt.Sprintf("Hello, %s! I'll help you track your budget!", username)
	bot.sendReply(update, text)
}

// Команда /help
func (bot *BudgetBot) commandHelp(update *tgbotapi.Update) {
	text := "Available commands:\n/start - Welcome message\n/help - Show this help"
	bot.sendReply(update, text)
}

// Неизвестная команда
func (bot *BudgetBot) commandUnknown(update *tgbotapi.Update) {
	text := "Sorry, I don't know this command. Use /help"
	bot.sendReply(update, text)
}

// Отправка ответа
func (bot *BudgetBot) sendReply(update *tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.api.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
