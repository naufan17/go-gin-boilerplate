package config

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"net/http"
	"time"
)

func SetupRateLimit(router *gin.Engine) {
	limiter := rate.NewLimiter(rate.Every(10*time.Minute), 100)

	router.Use(func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, please try again later",
			})
			c.Abort()
			return
		}
		c.Next()
	})
}
