package main

import (
	"github.com/naufan17/go-gin-boilerplate/config"
	routes "github.com/naufan17/go-gin-boilerplate/route"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	env := cfg.GinMode
	port := cfg.Port
	router := gin.Default()

	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// config.SetupSecure(router)
	config.SetupHelmet(router)
	config.SetupCORS(router)
	config.SetupRateLimit(router)
	// config.SetupCompress(router)

	routes.ApiRoutes(router)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Run(":" + port)
}
