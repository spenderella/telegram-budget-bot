package api

import (
	"fmt"
	"log"
	"strings"

	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *BudgetBot) handleMessage(update *tgbotapi.Update) {
	parts := strings.Fields(update.Message.Text)
	command := strings.ToLower(parts[0])

	switch command {
	case "/start":
		bot.commandStart(update)
	case "/help":
		bot.commandHelp(update)
	case "/add_expense":
		bot.commandAddExpense(update)
	default:
		bot.commandUnknown(update)
	}
}

func (bot *BudgetBot) commandStart(update *tgbotapi.Update) {
	username := update.Message.From.UserName
	if username == "" {
		username = update.Message.From.FirstName
	}

	text := fmt.Sprintf(StartMessage, username)
	bot.sendReply(update, text)
}

func (bot *BudgetBot) commandHelp(update *tgbotapi.Update) {
	text := HelpMessage
	bot.sendReply(update, text)
}

func (bot *BudgetBot) commandAddExpense(update *tgbotapi.Update) {

	userTime := time.Unix(int64(update.Message.Date), 0)
	userID := update.Message.From.ID

	amount, category, err := parseAddExpenseCommand(update.Message.Text)
	if err != nil {
		bot.sendReply(update, err.Error())
		return
	}

	err = bot.expenseService.AddExpense(userID, amount, category, userTime)
	if err != nil {
		bot.sendReply(update, "Failed to save expense. Please try again later.")
		return
	}

	bot.sendReply(update, "Expense saved successfully!")

}

func (bot *BudgetBot) commandUnknown(update *tgbotapi.Update) {
	text := "Sorry, I don't know this command. Use /help"
	bot.sendReply(update, text)
}

func (bot *BudgetBot) sendReply(update *tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.api.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
