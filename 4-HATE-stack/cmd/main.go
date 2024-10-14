package main

import (
	"awesome/handler"
	"awesome/service"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")

	articleService := service.ArticleService{}

	basicHandler := handler.BasicHandler{ArticleService: articleService}
	articleHandler := handler.ArticleHandler{ArticleService: articleService}

	app.GET("/", basicHandler.ShowHome)
	app.GET("/write", basicHandler.ShowWrite)
	app.GET("/articles/:id", articleHandler.ShowArticle)
	app.GET("/articles", articleHandler.CreateArticle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Logger.Fatal(app.Start(":" + port))

}
