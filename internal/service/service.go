package service

import (
	"context"
	"wikivin/internal/model"
)

type Services struct {
	Articles Articles
}

type Articles interface {
	LoadArticlesBriefInfo(ctx context.Context) ([]model.ArticleBriefInfo, error)
	LoadArticle(ctx context.Context, articleID int) (*model.ArticlePage, error)
}
