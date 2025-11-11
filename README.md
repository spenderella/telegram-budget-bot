# Telegram Budget Bot

WIP

Personal finance tracking bot for Telegram built with Go. 

## Features

- Add expense: amount and category
- Get your expenses. Optional filter by time period: today, week or month
- Get all existing categories
- Get expenses statistic by categories

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
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
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_USER="postgres"
export DB_PASSWORD="postgres"
export DB_NAME="budget_db"
export DB_SSLMODE="disable"
```
3. Start PostgreSQL database with Docker:
```bash
docker-compose up -d
```

Verify the database is running:
```bash
docker-compose ps
```
4. Run the bot:
```bash
go run cmd/budgetbot/main.go
```

To stop the database:
```bash
docker-compose down
```

## Database

The bot uses PostgreSQL for storing expenses and categories. Database connection parameters are configured via environment variables.

**Required environment variables:**
- `DB_HOST` - PostgreSQL host (default: localhost)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - PostgreSQL user
- `DB_PASSWORD` - PostgreSQL password
- `DB_NAME` - Database name
- `DB_SSLMODE` - SSL mode (usually "disable" for local development)

**Schema:**
- `expenses` 
- `categories` 
- `users`
- `schema_migration`


## Usage

Available commands:

/start - Welcome message
/help - Show available commands
/add_expense <amount> <category> - Add new expense with amount and category
/get_expenses - Get your expenses (limit 50). Optional: today, week, month
/get_categories - Get all categories
/get_statistics - Get total by categories (limit 50). Optional: today, week, month

Examples: 
/add_expense 25.50 coffee
/get_expenses today
/get_statistics month


