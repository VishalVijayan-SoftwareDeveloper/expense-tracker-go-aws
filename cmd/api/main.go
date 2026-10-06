package main

import (
	"log"

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

	log.Println("Database connected")

	r := router.Setup()

	log.Printf("Server started on %s", cfg.AppPort)

	err = r.Run(":" + cfg.AppPort)

	if err != nil {
		log.Fatal(err)
	}
}
