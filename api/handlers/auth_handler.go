package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/naufan17/go-gin-boilerplate/api/dtos"
	"github.com/naufan17/go-gin-boilerplate/config"
	"github.com/naufan17/go-gin-boilerplate/internal/services"
	"github.com/naufan17/go-gin-boilerplate/pkg/utils"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Register(c *gin.Context) {
	var user dtos.RegisterDto

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(user); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	_, err := services.RegisterUser(user)

	if err != nil {
		if err.Error() == "conflict" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User already exists",
			})

			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to register user",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
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
			"error": "Invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(user); validatorErr != nil {
		errors := utils.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	accessToken, err := services.LoginUser(user)

	if err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Email or password is incorrect",
			})

			return
		} else if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})

			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to login user",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to login user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": accessToken,
	})
}
