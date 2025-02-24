package handlers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/naufan17/go-gin-boilerplate/config"
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/services"
	"github.com/naufan17/go-gin-boilerplate/pkg/util"

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
		errors := util.ParseValidationError(validatorErr.(validator.ValidationErrors))

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
	var ipAddress string = c.ClientIP()
	var userAgent string = c.Request.UserAgent()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	if validatorErr := config.GetValidator().Struct(user); validatorErr != nil {
		errors := util.ParseValidationError(validatorErr.(validator.ValidationErrors))

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})

		return
	}

	accessToken, refreshToken, err := services.LoginUser(user, ipAddress, userAgent)

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

	expirationTime := time.Unix(int64(refreshToken.ExpiresIn), 0)
	c.SetCookie("refresh_token", refreshToken.RefreshToken, int(time.Until(expirationTime).Seconds()), "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"access_token": accessToken.AccessToken,
			"token_type":   accessToken.TokenType,
			"expires_in":   accessToken.ExpiresIn,
		},
	})
}
