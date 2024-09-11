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

func (r *ChaptersRepository) Create(ctx context.Context, chapter model.Chapter) (int, error) {

    query := "INSERT INTO chapters(article_id, parent_id, name, content) VALUES(?, ?, ?, ?)"
    

    result, err := r.db.ExecContext(ctx, query, chapter.ArticleID, chapter.ParentID, chapter.Name, chapter.Content)
    if err != nil {
        return -1, err
    }


    id, err := result.LastInsertId()
    if err != nil {
        return -1, err
    }

    return int(id), nil
}


func(r *ChaptersRepository)GetChaptersByArticleID(ctx context.Context, articleID int) ([]model.Chapter, error){
	var chapters []model.Chapter
	query:= "SELECT chapter_id, article_id, parent_id, name, content FROM chapters WHERE article_id = ?"
	err:= r.db.Select(&chapters, query, articleID)
	if err!= nil{
		return chapters, err
	}
	return chapters, nil
}