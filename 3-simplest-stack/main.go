package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "Hugo",
		})
	})

	e.Run(":8080")
}
