package expense

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) Profile(
	c *gin.Context,
) {

	userID :=
		c.GetString("user_id")

	c.JSON(
		http.StatusOK,
		gin.H{
			"user_id": userID,
		},
	)
}

func (h *Handler) CreateExpense(
	c *gin.Context,
) {

	userID :=
		c.GetString("user_id")

	var req CreateExpenseRequest

	if err :=
		c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.service.CreateExpense(
		c.Request.Context(),
		userID,
		req,
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "expense created",
		},
	)
}

func (h *Handler) GetExpenses(
	c *gin.Context,
) {

	userID :=
		c.GetString("user_id")

	expenses, err :=
		h.service.GetExpenses(
			c.Request.Context(),
			userID,
		)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		expenses,
	)
}
