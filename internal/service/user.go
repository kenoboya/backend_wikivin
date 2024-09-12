package service

import (
	"context"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
)

type UserService struct {
	userRepo repo.Users
}

func NewUserService(userRepo repo.Users) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) SignUp(ctx context.Context, requestSignUp model.RequestSignUp) (model.Token, error) {

}