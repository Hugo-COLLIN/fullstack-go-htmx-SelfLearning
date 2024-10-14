package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	InitDatabase()
	defer DB.Close() // defer enable to execute instruction when the function exits - here close the database connection

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(c *gin.Context) {
		todos := ReadToDoList()
		c.HTML(http.StatusOK, "task.html", gin.H{
			"todos": todos,
		})
	})

	e.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		status := c.PostForm("status")
		id, _ := CreateToDo(title, status)

		c.HTML(http.StatusOK, "task.html", gin.H{
			"id":     id,
			"title":  title,
			"status": status,
		})
	})

	e.DELETE("/todos/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteToDo(id)
	})

	e.Run(":8080")
}
