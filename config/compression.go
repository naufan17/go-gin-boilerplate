package config

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupCompress(router *gin.Engine) {
	router.Use(gzip.Gzip(gzip.DefaultCompression))
}
