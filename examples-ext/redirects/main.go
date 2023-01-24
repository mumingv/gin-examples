package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// 301-MovedPermanently，跳转到外部网站
	r.GET("/test1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com.hk")
	})
	// 302-Found，重定向到其他路由
	r.GET("/test2", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/index")
	})
	// 路由重定向
	r.GET("/test3", func(c *gin.Context) {
		c.Request.URL.Path = "/hello"
		r.HandleContext(c)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	r.Run(":8080")
}
