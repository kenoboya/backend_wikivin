package repo

import (
	"context"
	"wikivin/internal/model"
)

type Repositories struct {
	Articles Articles
	Chapters Chapters
	InfoBox InfoBox
}

type Articles interface {
	GetArticlesBriefInfo(ctx context.Context) ([]model.ArticleBriefInfo, error)
	GetArticleByID(ctx context.Context, articleID int) (model.Article, error)
}

type Chapters interface {
	GetChaptersByArticleID(ctx context.Context, articleID int) ([]model.Chapter, error)
}

type InfoBox interface{
	GetTypeAndObjectInfoBoxByArticleID(ctx context.Context, articleID int) (string, int, error)
	GetInfoBoxByObjectInfoBoxIDAndType(ctx context.Context, ObjectInfoBoxID int, infoBoxType string) (model.InfoBox, error)
}