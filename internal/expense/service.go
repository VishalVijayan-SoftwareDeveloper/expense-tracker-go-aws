package expense

import (
	"context"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(
	repo Repository,
) *Service {

	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateExpense(
	ctx context.Context,
	userID string,
	req CreateExpenseRequest,
) error {

	expenseDate, err := time.Parse(
		"2006-01-02",
		req.ExpenseDate,
	)

	if err != nil {
		return err
	}

	expense := &Expense{
		UserID:      userID,
		CategoryID:  req.CategoryID,
		Title:       req.Title,
		Description: req.Description,
		Amount:      req.Amount,
		ExpenseDate: expenseDate,
	}

	return s.repo.Create(
		ctx,
		expense,
	)
}
func (s *Service) GetExpenses(
	ctx context.Context,
	userID string,
) ([]Expense, error) {

	return s.repo.GetByUserID(
		ctx,
		userID,
	)
}
