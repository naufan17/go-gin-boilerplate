package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/go-gin-boilerplate/internal/controllers"
	"github.com/naufan17/go-gin-boilerplate/internal/middewares"
)

func SetupRouter(router *gin.Engine) {
	api := router.Group("/api/v1")

	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// User routes
		user := api.Group("/user")
		{
			user.GET("/profile", middewares.AuthenticationMiddleware, controllers.GetProfile)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Go Gin Rest API!",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Route not found",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "Method not allowed",
		})
	})
}
