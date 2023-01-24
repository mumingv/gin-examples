package main

import (
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/panic", func(c *gin.Context) {
		panic("panic-test-ok")
	})
	endless.ListenAndServe(":8080", router)
}
