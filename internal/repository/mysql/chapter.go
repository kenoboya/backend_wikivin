package repo

import (
	"context"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type ChaptersRepository struct {
	db *sqlx.DB
}

func NewChaptersRepository(db *sqlx.DB) *ChaptersRepository{
	return &ChaptersRepository{db}
}

func(r *ChaptersRepository) Create(ctx context.Context, chapter model.Chapter) error {
	query:= "INSERT INTO chapters(article_id, parent_id, name, content) VALUES($1, $2, $3)"
	_, err:= r.db.Exec(query, chapter.ArticleID, chapter.ParentID, chapter.Name, chapter.Content)
	if err != nil{
		return err
	}
	return nil
}

func(r *ChaptersRepository)GetChaptersByArticleID(ctx context.Context, articleID int) ([]model.Chapter, error){
	var chapters []model.Chapter
	query:= "SELECT chapter_id, article_id, parent_id, name, content FROM chapters WHERE article_id = $1"
	err:= r.db.Select(&chapters, query, articleID)
	if err!= nil{
		return chapters, err
	}
	return chapters, nil
}