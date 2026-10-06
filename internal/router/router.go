package router

import (
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/health"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	r.GET("/health", health.Health)

	return r
}
