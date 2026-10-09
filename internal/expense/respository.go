package expense

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(
		ctx context.Context,
		expense *Expense,
	) error

	GetByUserID(
		ctx context.Context,
		userID string,
	) ([]Expense, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(
	db *pgxpool.Pool,
) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(
	ctx context.Context,
	expense *Expense,
) error {

	query := `
	INSERT INTO expenses
	(
		user_id,
		category_id,
		title,
		description,
		amount,
		expense_date
	)
	VALUES
	(
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		expense.UserID,
		expense.CategoryID,
		expense.Title,
		expense.Description,
		expense.Amount,
		expense.ExpenseDate,
	)

	return err
}

func (r *repository) GetByUserID(
	ctx context.Context,
	userID string,
) ([]Expense, error) {

	query := `
	SELECT
		id,
		user_id,
		category_id,
		title,
		description,
		amount,
		expense_date,
		created_at,
		updated_at
	FROM expenses
	WHERE user_id = $1
	ORDER BY expense_date DESC
	`

	rows, err :=
		r.db.Query(
			ctx,
			query,
			userID,
		)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var expenses []Expense

	for rows.Next() {

		var expense Expense

		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.CategoryID,
			&expense.Title,
			&expense.Description,
			&expense.Amount,
			&expense.ExpenseDate,
			&expense.CreatedAt,
			&expense.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		expenses = append(
			expenses,
			expense,
		)
	}

	return expenses, nil
}
