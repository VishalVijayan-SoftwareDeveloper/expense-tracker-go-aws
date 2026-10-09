package router

import (
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/auth"
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/expense"
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(
	authHandler *auth.Handler,
	expenseHandler *expense.Handler,
	authMiddleware *middleware.AuthMiddleware,
) *gin.Engine {

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	api := r.Group("/api/v1")

	// Public routes
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// Protected routes
	protected := api.Group("/")
	protected.Use(authMiddleware.Authenticate())
	{
		protected.GET("/profile", expenseHandler.Profile)
		protected.POST(
			"/expenses",
			expenseHandler.CreateExpense,
		)

		protected.GET(
			"/expenses",
			expenseHandler.GetExpenses,
		)
	}

	return r
}
