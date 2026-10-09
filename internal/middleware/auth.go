package middleware

import (
	"github.com/VishalVijayan-SoftwareDeveloper/expense-tracker-go-aws/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	secret string
}

func NewAuthMiddleware(
	secret string,
) *AuthMiddleware {

	return &AuthMiddleware{
		secret: secret,
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader :=
			c.GetHeader("Authorization")

		if authHeader == "" {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing token",
				},
			)

			c.Abort()

			return
		}

		parts := strings.Split(
			authHeader,
			" ",
		)

		if len(parts) != 2 {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			c.Abort()

			return
		}

		token := parts[1]

		userID, err :=
			auth.ValidateToken(
				token,
				m.secret,
			)

		if err != nil {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			c.Abort()

			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}
