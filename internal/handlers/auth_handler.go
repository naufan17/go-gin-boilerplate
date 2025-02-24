package handlers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/naufan17/go-gin-boilerplate/config"
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/services"
	"github.com/naufan17/go-gin-boilerplate/pkg/auth"
	"github.com/naufan17/go-gin-boilerplate/pkg/util"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Register(c *gin.Context) {
	var user dtos.RegisterDto

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
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
				"error": "user already exists",
			})

			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to register user",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to register user",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
	})
}

func Login(c *gin.Context) {
	var user dtos.LoginDto
	var ipAddress string = c.ClientIP()
	var userAgent string = c.Request.UserAgent()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
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
				"error": "email or password is incorrect",
			})

			return
		} else if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})

			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to login user",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to login user",
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

func RefreshToken(c *gin.Context) {
	claimsSession := c.MustGet("claimsSession").(*auth.Claims)
	id := claimsSession.Sub

	accessToken, err := services.RefreshTokenUser(id)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})

			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to refresh token",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to refresh token",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"access_token": accessToken.AccessToken,
			"token_type":   accessToken.TokenType,
			"expires_in":   accessToken.ExpiresIn,
		},
	})
}

func Logout(c *gin.Context) {
	claimsUser := c.MustGet("claimsUser").(*auth.Claims)
	id := claimsUser.Sub

	err := services.LogoutUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to logout user",
		})

		return
	}

	c.SetCookie("refresh_token", "", -1, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "user logged out successfully",
	})
}
