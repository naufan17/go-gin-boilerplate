package config

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func SetupSecure(router *gin.Engine) {
	router.Use(secure.New(secure.Config{
		// SSLRedirect:           true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		ContentSecurityPolicy: "default-src 'self'",
	}))
}
