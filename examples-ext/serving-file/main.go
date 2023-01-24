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
		// File()将文件内容直接返回，Content-Type: text/plain
		c.File("local/file.go")
	})

	var fs http.FileSystem = http.Dir("my_file_system")
	router.GET("/fs/file", func(c *gin.Context) {
		// FileFromFS()与相比，支持指定http.FileSystem
		c.FileFromFS("fs/file.go", fs)
	})

	/// Serving data from reader

	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("http://httpbin.org/image/png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		defer reader.Close()
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.Run(":8080")
}
