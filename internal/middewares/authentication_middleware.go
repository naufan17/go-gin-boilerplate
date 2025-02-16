package middewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/go-gin-boilerplate/pkg/auth"
)

func AuthenticateJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is required",
		})

		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := auth.ValidateJWT(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})

		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}
