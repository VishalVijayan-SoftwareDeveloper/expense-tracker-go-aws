package health

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	DB *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) Health(c *gin.Context) {

	err := h.DB.Ping(context.Background())

	if err != nil {

		c.JSON(500, gin.H{
			"status": "DOWN",
		})

		return
	}

	c.JSON(200, gin.H{
		"status":   "UP",
		"database": "CONNECTED",
	})
}
