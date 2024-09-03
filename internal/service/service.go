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
	LoadArticlesBriefInfo(ctx context.Context) ([]model.ArticleBriefInfo, error)
	LoadArticle(ctx context.Context, title string) (*model.ArticlePage, error)
}
