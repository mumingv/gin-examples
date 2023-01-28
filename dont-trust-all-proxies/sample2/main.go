package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// Use predefined header gin.PlatformXXX
	router.TrustedPlatform = gin.PlatformGoogleAppEngine
	// Or set your own trusted request header for another trusted proxy service
	// Don't set it to any suspect request header, it's unsafe
	router.TrustedPlatform = "X-CDN-IP"

	router.GET("/", func(c *gin.Context) {
		// If you set TrustedPlatform, ClientIP() will resolve the
		// corresponding header and return IP directly
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
	})
	router.Run()
}
