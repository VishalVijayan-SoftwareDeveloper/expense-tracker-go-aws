package router

import (
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/auth"

	"github.com/gin-gonic/gin"
)

func Setup(
	authHandler *auth.Handler,
) *gin.Engine {

	r := gin.Default()

	api := r.Group("/api/v1")

	{
		api.POST(
			"/auth/register",
			authHandler.Register,
		)

		api.POST(
			"/auth/login",
			authHandler.Login,
		)
	}

	r.GET("/health", func(c *gin.Context) {

		c.JSON(
			200,
			gin.H{
				"status": "UP",
			},
		)
	})

	return r
}
