package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(router *gin.Engine) {
	allowOrigin := LoadConfig().AllowedOrigin

	if allowOrigin == "" {
		allowOrigin = "*"
	}

	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))
}
