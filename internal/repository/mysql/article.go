package repo

import (
	"context"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type ArticlesRepository struct {
	db *sqlx.DB
}

func NewArticlesRepository (db *sqlx.DB) *ArticlesRepository{
	return &ArticlesRepository{db}
}

func (r *ArticlesRepository) Create(ctx context.Context, article model.Article) (int, error) {
    query := `INSERT INTO articles (title, lead_section, image) VALUES (?, ?, ?)`

    result, err := r.db.ExecContext(ctx, query, article.Title, article.LeadSection, article.Image)
    if err != nil {
        return -1, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return -1, err
    }

    return int(id), nil
}

func (r *ArticlesRepository)GetArticles(ctx context.Context) ([]model.Article, error){
	var articles []model.Article
	query:= "SELECT article_id,title, lead_section, image FROM articles"
	err:= r.db.Select(&articles, query)
	if err!= nil{
		return articles, err
	}
	return articles, nil
}	
func(r *ArticlesRepository)GetArticleByTitle(ctx context.Context, title string) (model.Article, error){
	var article model.Article
	err:= r.db.Get(&article, "SELECT * FROM articles WHERE title = ?", title)
	if err != nil{
		return article, err
	}
	return article, nil
}