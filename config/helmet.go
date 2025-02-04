package config

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
)

func SetupHelmet(router *gin.Engine) {
	router.Use(helmet.Default())
}
