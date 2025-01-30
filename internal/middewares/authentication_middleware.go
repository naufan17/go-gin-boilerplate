package middewares

import (
	"net/http"
	"strings"

	"github.com/naufan17/go-gin-boilerplate/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ValidateJWT(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}
