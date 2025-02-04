package main

import (
	"github.com/naufan17/go-gin-boilerplate/config"
	"github.com/naufan17/go-gin-boilerplate/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// config.SetupSecure(router)
	config.SetupHelmet(router)
	config.SetupCORS(router)
	config.SetupRateLimit(router)
	// config.SetupCompress(router)
	routes.ApiRoutes(router)

	cfg := config.LoadConfig()
	port := cfg.Port

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
