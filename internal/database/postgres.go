package database

import (
	"context"
	"fmt"

	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) (*pgxpool.Pool, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	return db, nil
}
