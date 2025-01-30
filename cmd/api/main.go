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

	// Initialize the database
	// db := configs.ConnectDB()

	// Migrate the database
	// configs.MigrateDB(db)

	// Seed the database
	// seeders.SeedAll(db)

	// Initialize the router
	router := gin.Default()

	// Setup the CORS middleware
	configs.SetupCORS(router)

	// Setup the secure middleware
	configs.SetupSecure(router)

	// Setup the routes
	routes.SetupRouter(router)

	// Get the port from environment variable
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// Start the server
	router.Run(":" + port)
}
