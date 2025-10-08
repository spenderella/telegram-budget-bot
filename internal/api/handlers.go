package api

import (
	"fmt"
	"log"
	"strings"

	//"telegram-finance-bot/internal/models"
	"telegram-finance-bot/internal/constants"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bot *BudgetBot) handleMessage(update *tgbotapi.Update) {
	parts := strings.Fields(update.Message.Text)
	command := strings.ToLower(parts[0])

	switch command {
	case constants.CmdStart:
		bot.commandStart(update)
	case constants.CmdHelp:
		bot.commandHelp(update)
	case constants.CmdAddExpense:
		bot.commandAddExpense(update)
	case constants.CmdGetExpenses:
		bot.commandGetExpenses(update)
	case constants.CmdGetCategories:
		bot.commandGetCategories(update)
	default:
		bot.commandUnknown(update)
	}
}

func (bot *BudgetBot) commandStart(update *tgbotapi.Update) {
	username := update.Message.From.UserName
	if username == "" {
		username = update.Message.From.FirstName
	}

	text := fmt.Sprintf(constants.StartMessage, username)
	bot.sendReply(update, text)
}

func (bot *BudgetBot) commandHelp(update *tgbotapi.Update) {
	text := constants.HelpMessage
	bot.sendReply(update, text)
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
