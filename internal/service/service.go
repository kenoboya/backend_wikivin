package service

import (
	"context"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
)

type Services struct {
	Articles Articles
}

func NewServices(repo *repo.Repositories)*Services{
	return &Services{
		Articles: NewArticlesService(repo.Articles, repo.Chapters, repo.InfoBox),
	}
}

type Articles interface {
	CreateArticle(ctx context.Context, infoBoxDB model.InfoBoxDB,article model.Article) error
	LoadArticles(ctx context.Context) ([]model.Article, error)
	LoadArticle(ctx context.Context, title string) (*model.ArticlePage, error)
}
