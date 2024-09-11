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
	Create(ctx context.Context, article model.Article) (int, error)
	GetArticles(ctx context.Context) ([]model.Article, error)
	GetArticleByTitle(ctx context.Context, title string) (model.Article, error)
}

type Chapters interface {
	Create(ctx context.Context, chapter model.Chapter) error 
	GetChaptersByArticleID(ctx context.Context, articleID int) ([]model.Chapter, error)
}

type InfoBox interface{
	Create(ctx context.Context, articleID int, infoBoxID int) error
	CreateInfoBoxByType(ctx context.Context, infoBoxDB model.InfoBoxDB) (int, error)
	GetTypeAndObjectInfoBoxByArticleID(ctx context.Context, articleID int) (string, int, error)
	GetInfoBoxByObjectInfoBoxIDAndType(ctx context.Context, ObjectInfoBoxID int, infoBoxType string) (model.InfoBox, error)
}