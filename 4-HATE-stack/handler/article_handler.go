package handler

import (
	"awesome/service"
	"awesome/view/article"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	ArticleService service.ArticleService
}

func (h *ArticleHandler) ShowArticle(c echo.Context) error {
	id := c.Param("id")
	article := h.ArticleService.Read(id)
	return render(c, article.Show(article))
}

func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	t := c.FormValue("title")
	b := c.FormValue("content")
	a := h.ArticleService.Create(t, b)
	return render(c, article.Show(a))
}
