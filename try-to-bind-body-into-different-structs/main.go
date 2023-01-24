package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ParamsOne struct {
	Username string `json:"username"`
}

type ParamsTwo struct {
	Username string `json:"username"`
}

func main() {
	r := gin.New()
	r.POST("/", func(c *gin.Context) {
		var f ParamsOne
		if err := c.ShouldBindBodyWith(&f, binding.JSON); err != nil {
			log.Printf("%+v", err)
		}
		log.Printf("%+v", f)

		var ff ParamsTwo
		if err := c.ShouldBindBodyWith(&ff, binding.JSON); err != nil {
			log.Printf("%+v", err)
		}
		log.Printf("%+v", ff)
		c.IndentedJSON(http.StatusOK, f)
	})
	r.Run(":8080")
}
