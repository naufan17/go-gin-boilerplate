package config

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"net/http"
	"time"
)

func SetupRateLimit(router *gin.Engine) {
	cfg := LoadConfig()
	maxRequests, err := strconv.Atoi(cfg.MaxRequests)

	if err != nil {
		panic("Invalid MaxRequests value")
	}

	windowTime, err := time.ParseDuration(cfg.WindowTime + "ms")

	if err != nil {
		panic("Invalid WindowTime value: " + cfg.WindowTime)
	}

	limiter := rate.NewLimiter(rate.Every(windowTime), maxRequests)

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
