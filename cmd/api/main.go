package main

import (
	"github.com/naufan17/go-gin-boilerplate/internal/configs"
	"github.com/naufan17/go-gin-boilerplate/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// configs.SetupSecure(router)
	configs.SetupHelmet(router)
	configs.SetupCORS(router)
	configs.RateLimit(router)
	// configs.SetupCompress(router)
	routes.SetupRouter(router)

	cfg := configs.LoadConfig()
	port := cfg.Port

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
