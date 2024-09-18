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
		a.title,
		a.description, 
		a.image AS article_image
	FROM articles a
	INNER JOIN favorite_articles f ON f.article_id = a.article_id
	INNER JOIN people p ON p.person_id = a.author_id
	WHERE f.user_id = ?`

	err:= r.db.Select(&favoriteArticles, query,  userID)
	if err != nil{
		return []model.FavoriteArticle{}, err
	}
	return favoriteArticles, nil
}

func (r *FavoriteRepository) AddFavoriteByUserAndArticleID(ctx context.Context, userID int, articleID int)error{
	query:= "INSERT INTO favorite_articles(user_id, article_id) VALUES(?,?)"
	if _, err:= r.db.ExecContext(ctx,query,userID, articleID); err != nil{
		return err
	}
	return nil
}

func (r *FavoriteRepository) DeleteFavoriteByUserAndArticleID(ctx context.Context, userID int, articleID int)error{
	query:= "DELETE FROM favorite_articles WHERE user_id = ? AND article_id = ?"
	if _, err:= r.db.ExecContext(ctx,query, userID, articleID); err != nil{
		return err
	}
	return nil
}
