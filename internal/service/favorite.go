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