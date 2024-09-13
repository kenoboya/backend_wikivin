package service

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
	"wikivin/pkg/auth"
	"wikivin/pkg/hash"
)

type UsersService struct {
	userRepo repo.Users
	peopleRepo repo.People
	hasher hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL time.Duration
	refreshTokenTTL time.Duration
}

func NewUsersService(userRepo repo.Users, peopleRepo repo.People, 
	hasher hash.PasswordHasher, tokenManager auth.TokenManager) *UsersService {
	return &UsersService{
		userRepo: userRepo,
		peopleRepo: peopleRepo,
		hasher: hasher,
		tokenManager: tokenManager,
	}
}

func (s *UsersService) SignUp(ctx context.Context, requestSignUp model.RequestSignUp) (model.Tokens, error) {
	passwordHash, err := s.hasher.Hash(requestSignUp.UserSignUp.Password)
	if err != nil{
		return model.Tokens{}, err
	}
	requestSignUp.UserSignUp.Password = passwordHash
	user, err := s.userRepo.Create(ctx, requestSignUp.UserSignUp)
	if err != nil{
		return model.Tokens{}, err
	}
	requestSignUp.Person.UserID = user.ID
	if err:= s.peopleRepo.Create(ctx, requestSignUp.Person); err != nil{
		return model.Tokens{}, err
	}
	return s.createSession(ctx, &user)
}

func (s *UsersService) SignIn(ctx context.Context, requestSignIn model.UserSignIn)(model.Tokens, error){
	user, err:= s.userRepo.GetByLogin(ctx, requestSignIn.Login)
	if err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return model.Tokens{}, model.ErrInvalidLogin
		}
		return model.Tokens{}, err
	}
	if user.Blocked{
		return model.Tokens{}, model.ErrUserBlocked
	}
	if err:= s.hasher.Compare(user.Password, requestSignIn.Password); err != nil{
		return model.Tokens{}, err
	}
	return s.createSession(ctx, &user)
}

func (s *UsersService) createSession(ctx context.Context, user *model.User) (model.Tokens, error){
	var (
		res model.Tokens
		err error
	)	
	res.AccessToken, err = s.tokenManager.NewJWT(user.ID, user.Role, s.accessTokenTTL)
	if err != nil{
		return model.Tokens{}, err
	}
	res.RefreshToken, err = s.tokenManager.NewJWT(user.ID, user.Role, s.refreshTokenTTL)
	if err != nil{
		return model.Tokens{}, err
	}
	return res, nil
}