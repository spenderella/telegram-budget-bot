package repositories

import (
	"database/sql"
	//"log"

	"telegram-finance-bot/internal/models"
)

type ExpenseRepository struct {
	//Expenses map[int64][]models.Expense //deprecate
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{
		//Expenses: make(map[int64][]models.Expense),
		db: db,
	}
}

func (r *ExpenseRepository) Save(expense models.Expense) error {
	var expenseId int
	query := `
        INSERT INTO expenses (user_id, category_id, amount, currency) 
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	err := r.db.QueryRow(
		query, expense.UserID, expense.Category.ID, expense.Amount, expense.Currency).Scan(
		&expenseId)

	return err

}

func (r *ExpenseRepository) GetExpenses(filter models.ExpenseFilter) ([]models.Expense, error) {

	query := `
        SELECT e.user_id, c.id, c.name, e.amount, e.currency, e.created_at 
		FROM expenses e
		LEFT JOIN users u ON e.user_id = u.id 
		LEFT JOIN categories c ON e.category_id = c.id
		WHERE u.telegram_id = $1
			AND e.created_at >= $2
          	AND e.created_at <= $3
		ORDER BY e.created_at DESC
		LIMIT $4
    `

	rows, err := r.db.Query(query, filter.UserTgID, filter.DateFrom, filter.DateTo, filter.Limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense

	for rows.Next() {
		var exp models.Expense
		exp.Category = &models.Category{}

		err := rows.Scan(
			&exp.UserID,
			&exp.Category.ID,
			&exp.Category.Name,
			&exp.Amount,
			&exp.Currency,
			&exp.Date,
		)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, exp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}
