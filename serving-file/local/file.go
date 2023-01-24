package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	/// Serving static files

	// Static()不显示文件
	router.Static("/assets", "./assets")
	// StaticFS()显示文件
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	// StaticFile()单个文件
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// StaticFileFS()单个文件，指定目录
	router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))

	/// Serving data from file

	router.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})

	router.Run(":8080")
}
