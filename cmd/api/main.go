package main

import (
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/expense"
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/middleware"
	"log"

	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/auth"
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/config"
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/database"
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/router"
)

func main() {

	cfg := config.Load()

	db, err := database.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	authRepo := auth.NewRepository(db)

	authService := auth.NewService(
		authRepo,
		cfg,
	)

	authHandler := auth.NewHandler(
		authService,
	)
	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(
		cfg.JWTSecret,
	)

	expenseRepo := expense.NewRepository(db)

	expenseService := expense.NewService(
		expenseRepo,
	)

	expenseHandler := expense.NewHandler(
		expenseService,
	)

	r := router.Setup(
		authHandler,
		expenseHandler,
		authMiddleware,
	)

	log.Printf(
		"Server started on port %s",
		cfg.AppPort,
	)

	log.Fatal(
		r.Run(":" + cfg.AppPort),
	)
}
