package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/naufan17/go-gin-boilerplate/internal/services"
	"github.com/naufan17/go-gin-boilerplate/internal/utils"

	"net/http"
)

func GetProfile(c *gin.Context) {
	claims := c.MustGet("claims").(*utils.Claims)
	id := claims.Sub
	user, err := services.ProfileUser(id)

	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
