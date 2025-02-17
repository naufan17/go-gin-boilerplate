package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/naufan17/go-gin-boilerplate/config"
	"github.com/naufan17/go-gin-boilerplate/internal/dtos"
	"github.com/naufan17/go-gin-boilerplate/internal/services"
	"github.com/naufan17/go-gin-boilerplate/pkg/auth"
	"github.com/naufan17/go-gin-boilerplate/pkg/util"

	"net/http"
)

func GetProfile(c *gin.Context) {
	claims := c.MustGet("claims").(*auth.Claims)
	id := claims.Sub
	user, err := services.ProfileUser(id)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user profile",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateProfile(c *gin.Context) {
	claims := c.MustGet("claims").(*auth.Claims)
	id := claims.Sub
	var user dtos.UpdateProfileDto

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

	_, err := services.UpdateProfileUser(user, id)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user profile",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User profile updated successfully",
	})
}

func UpdatePassword(c *gin.Context) {
	claims := c.MustGet("claims").(*auth.Claims)
	id := claims.Sub
	var user dtos.UpdatePasswordDto

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

	_, err := services.UpdatePasswordUser(user, id)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})

			return
		} else if err.Error() == "internal server error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update user password",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user password",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User password updated successfully",
	})
}
