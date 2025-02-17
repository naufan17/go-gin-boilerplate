package main

import (
	"github.com/naufan17/go-gin-boilerplate/config"
	routes "github.com/naufan17/go-gin-boilerplate/route"

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

	router.Run(":" + port)
}
