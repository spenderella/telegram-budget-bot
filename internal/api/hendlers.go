package api

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Обработчик команды /start
func handleStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	text := fmt.Sprintf("Hello, %s! This bot will count your expenses!", update.Message.From.UserName)
	sendReply(bot, update, text)
}

// Обработчик команды /help
func handleHelp(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	text := "Available commands:\n/start - Welcome message\n/help - Show this help"
	sendReply(bot, update, text)
}

// Обработчик неизвестных команд
func handleUnknown(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	text := "Sorry, I don't know this command. Use /help"
	sendReply(bot, update, text)
}

// Функция отправки ответа
func sendReply(bot *tgbotapi.BotAPI, update *tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
