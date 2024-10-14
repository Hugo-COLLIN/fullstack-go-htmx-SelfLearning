package service

type ArticleService struct{}

func (s *ArticleService) Create(title string, body string) model.Article {
	return Article{Title: title, Body: body}
}
