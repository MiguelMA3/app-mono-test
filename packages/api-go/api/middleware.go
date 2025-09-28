package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/api-go/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token de autorizacao necessario"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
