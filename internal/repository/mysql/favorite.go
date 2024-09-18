package repo

import (
	"context"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type FavoriteRepository struct {
	db *sqlx.DB
}

func NewFavoriteRepository (db *sqlx.DB)*FavoriteRepository{
	return &FavoriteRepository{db}
}

func(r *FavoriteRepository)GetFavoriteArticlesByUserID(ctx context.Context, userID int)([]model.FavoriteArticle, error){
	var favoriteArticles []model.FavoriteArticle
	query := `
	SELECT 
		a.author_id, 
		p.image AS author_image, 
		p.first_name, 
		p.last_name, 
		a.article_id,
		a.description, 
		a.image AS article_image
	FROM articles a
	INNER JOIN favorite_articles f ON f.article_id = a.article_id
	INNER JOIN people p ON p.author_id = a.author_id
	WHERE f.user_id = ?`

	err:= r.db.Select(&favoriteArticles, query,  userID)
	if err != nil{
		return []model.FavoriteArticle{}, err
	}
	return favoriteArticles, nil
}