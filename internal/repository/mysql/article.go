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
func (r *ArticlesRepository)GetArticlesBriefInfo(ctx context.Context) ([]model.ArticleBriefInfo, error){
	var articlesBriefInfo []model.ArticleBriefInfo
	query:= "SELECT title, lead_section, image FROM articles"
	err:= r.db.Select(&articlesBriefInfo, query)
	if err!= nil{
		return articlesBriefInfo, err
	}
	return articlesBriefInfo, nil
}	
func(r *ArticlesRepository)GetArticleByTitle(ctx context.Context, title string) (model.Article, error){
	var article model.Article
	err:= r.db.Get(&article, "SELECT * FROM articles WHERE title = $1", title)
	if err != nil{
		return article, err
	}
	return article, nil
}