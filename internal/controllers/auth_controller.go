package controllers

import (
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/services"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Register(c *gin.Context) {
	var user dtos.RegisterDto

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	_, err := services.RegisterUser(user)

	if err != nil {
		if err.Error() == "conflict" {
			c.JSON(http.StatusConflict, gin.H{
				"message": "User already exists",
			})
			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to register user",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {
	var user dtos.LoginDto

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	accessToken, err := services.LoginUser(user)

	if err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Email or password is incorrect",
			})
			return
		} else if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to login user",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to login user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": accessToken,
	})
}
