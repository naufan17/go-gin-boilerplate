package middewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/go-gin-boilerplate/pkg/auth"
)

func AuthorizeBearer(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header is required",
		})

		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := auth.ValidateJWTAccess(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})

		c.Abort()
		return
	}

	c.Set("claimsUser", claims)
	c.Next()
}

func AuthorizeCookie(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "refresh token cookie is required",
		})

		c.Abort()
		return
	}

	claims, err := auth.ValidateJWTRefresh(cookie)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid refresh token",
		})

		c.Abort()
		return
	}

	c.Set("claimsSession", claims)
	c.Next()
}
