# Telegram Budget Bot

WIP

Personal finance tracking bot for Telegram built with Go. 

## Features

- Add expense: amount and category
- Get your expenses

## Prerequisites

- Go 1.21+
- Telegram Bot Token (from @BotFather)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/spenderella/telegram-budget-bot.git
cd telegram-budget-bot 
```
2. Set up environment variable:
```bash
export BOT_TOKEN="your_bot_token_here"
```

3. Run the bot:
```bash
go run cmd/budgetbot/main.go
```

## Usage

Available commands:

/start - Welcome message
/help - Show available commands
/add_expense <amount> <category> - Add new expense with amount and category
/get_expenses - Get your expenses (limit 50). Optional: today, week, month

Examples: 
/add_expense 25.50 coffee
/get_expenses today


