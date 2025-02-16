package main

import (
	"github.com/naufan17/go-gin-boilerplate/config"
	"github.com/naufan17/go-gin-boilerplate/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadConfig().GinMode

	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	// config.SetupSecure(router)
	config.SetupHelmet(router)
	config.SetupCORS(router)
	config.SetupRateLimit(router)
	// config.SetupCompress(router)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.ApiRoutes(router)

	cfg := config.LoadConfig()
	port := cfg.Port

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
