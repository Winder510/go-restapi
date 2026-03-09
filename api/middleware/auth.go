package middleware

import (
	"net/http"
	"strings"

	"go_mysql/internal/model"
	"go_mysql/internal/util"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.Response{
				Success: false,
				Error:   "Authorization header required",
			})
			c.Abort()
			return
		}

		// Bearer token format: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, model.Response{
				Success: false,
				Error:   "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := util.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Response{
				Success: false,
				Error:   "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
