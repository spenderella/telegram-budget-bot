package models

type Category struct {
	ID   int
	Name string
}

type CategoryExpenses struct {
	Category      *Category
	Total         float64
	TotalCurrency string
	//Expenses []Expense
}
