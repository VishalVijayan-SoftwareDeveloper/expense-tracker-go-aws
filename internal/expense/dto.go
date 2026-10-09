package expense

type CreateExpenseRequest struct {
	CategoryID  string  `json:"category_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	ExpenseDate string  `json:"expense_date"`
}

type ExpenseResponse struct {
	ID          string  `json:"id"`
	CategoryID  string  `json:"category_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	ExpenseDate string  `json:"expense_date"`
}
