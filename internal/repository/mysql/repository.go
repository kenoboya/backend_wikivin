package repo

import (
	"context"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Articles Articles
	Chapters Chapters
	InfoBox InfoBox
}

func NewRepositories(db *sqlx.DB) *Repositories{
	return &Repositories{
		Articles: NewArticlesRepository(db),
		Chapters: NewChaptersRepository(db),
		InfoBox: NewInfoBoxesRepository(db),
	}
}

type Articles interface {
	GetArticlesBriefInfo(ctx context.Context) ([]model.ArticleBriefInfo, error)
	GetArticleByTitle(ctx context.Context, title string) (model.Article, error)
}

type Chapters interface {
	GetChaptersByArticleID(ctx context.Context, articleID int) ([]model.Chapter, error)
}

type InfoBox interface{
	GetTypeAndObjectInfoBoxByArticleID(ctx context.Context, articleID int) (string, int, error)
	GetInfoBoxByObjectInfoBoxIDAndType(ctx context.Context, ObjectInfoBoxID int, infoBoxType string) (model.InfoBox, error)
}