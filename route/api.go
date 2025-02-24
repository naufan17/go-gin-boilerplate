package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufan17/go-gin-boilerplate/internal/handlers"
	"github.com/naufan17/go-gin-boilerplate/internal/middewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/naufan17/go-gin-boilerplate/docs"
)

func ApiRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.GET("/refresh", middewares.AuthenticateCookie, handlers.RefreshToken)
			auth.GET("/logout", middewares.AuthenticateJWT, middewares.AuthenticateCookie, handlers.Logout)
		}

		account := api.Group("/account")
		{
			account.GET("/profile", middewares.AuthenticateJWT, handlers.GetProfile)
			account.GET("/session", middewares.AuthenticateJWT, handlers.GetSession)
			account.POST("/update-profile", middewares.AuthenticateJWT, handlers.UpdateProfile)
			account.POST("/update-password", middewares.AuthenticateJWT, handlers.UpdatePassword)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Go Gin Rest API!",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "method not allowed",
		})
	})
}
