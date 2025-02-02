package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/naufan17/go-gin-boilerplate/internal/configs"
	"github.com/naufan17/go-gin-boilerplate/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// configs.SetupSecure(router)
	configs.SetupHelmet(router)
	configs.SetupCORS(router)
	configs.RateLimit(router)
	routes.SetupRouter(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
