package service

import (
	"context"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
)

type FavoritesService struct {
	favoriteRepo repo.Favorites
}

func NewFavoritesService(favoriteRepo repo.Favorites) *FavoritesService {
	return &FavoritesService{favoriteRepo: favoriteRepo}
}

func (s *FavoritesService) GetFavoriteArticlesByUserID(ctx context.Context, userID int)([]model.FavoriteArticle, error){
	return s.favoriteRepo.GetFavoriteArticlesByUserID(ctx, userID)
}

func (s *FavoritesService) AddFavorite(ctx context.Context, userID int, articleID int) error{
	return s.favoriteRepo.AddFavoriteByUserAndArticleID(ctx, userID, articleID)
}

func (s *FavoritesService) DeleteFavorite(ctx context.Context, userID int, articleID int) error{
	return s.favoriteRepo.DeleteFavoriteByUserAndArticleID(ctx, userID, articleID)
}