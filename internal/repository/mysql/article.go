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

func (r *ArticlesRepository) Create(ctx context.Context,article model.Article) (int, error) {
    var id int

    query := `INSERT INTO articles (title, lead_section, image) VALUES (:title, :lead_section, :image) RETURNING article_id`
    row := r.db.QueryRow(query, article)

    if err := row.Scan(&id); err != nil {
        return -1, err
    }

    return id, nil
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
	err:= r.db.Get(&article, "SELECT * FROM articles WHERE title = $1", title)
	if err != nil{
		return article, err
	}
	return article, nil
}