package expense

import "time"

type Expense struct {
	ID          string
	UserID      string
	CategoryID  string
	Title       string
	Description string
	Amount      float64
	ExpenseDate time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
