package auth

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/user"
)

type Repository interface {
	CreateUser(ctx context.Context, user *user.User) error
	GetByEmail(ctx context.Context, email string) (*user.User, error)
}
type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(
	ctx context.Context,
	user *user.User,
) error {

	query := `
	INSERT INTO users
	(
		email,
		password_hash
	)
	VALUES
	(
		$1,
		$2
	)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
	)

	return err
}

func (r *repository) GetByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {

	var user user.User

	query := `
	SELECT
		id,
		email,
		password_hash,
		created_at,
		updated_at
	FROM users
	WHERE email = $1
	`

	err := r.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

var _ pgx.Row
