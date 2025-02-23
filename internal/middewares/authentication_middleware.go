package middewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/go-gin-boilerplate/pkg/auth"
)

func AuthenticateJWTAccess(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is required",
		})

		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := auth.ValidateJWTAccess(token)

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

func AuthenticateJWTRefresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Refresh token cookie is required",
		})

		c.Abort()
		return
	}

	claims, err := auth.ValidateJWTRefresh(cookie)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid refresh token",
		})

		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}
