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

func (s *UsersService) GetAccessTokenTTL() time.Duration{
	return s.accessTokenTTL
}

func (s *UsersService) GetRefreshTokenTTL() time.Duration{
	return s.refreshTokenTTL
}

func NewUsersService(userRepo repo.Users, peopleRepo repo.People, 
	hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *UsersService {
	return &UsersService{
		userRepo: userRepo,
		peopleRepo: peopleRepo,
		hasher: hasher,
		tokenManager: tokenManager,
		accessTokenTTL: accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
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
	if user.IsBlocked(){
		return model.Tokens{}, model.ErrUserBlocked
	}
	if err:= s.hasher.Compare(user.Password, requestSignIn.Password); err != nil{
		return model.Tokens{}, err
	}
	return s.createSession(ctx, &user)
}

func (s *UsersService) RefreshToken(ctx context.Context, refreshToken string) (model.Tokens, error){
	var (
		res model.Tokens
		err error
	)
	if res.RefreshToken, err = s.tokenManager.RefreshToken(refreshToken, s.refreshTokenTTL); err != nil{
		return model.Tokens{}, err
	}
	claims, err:= s.tokenManager.ParseToken(res.RefreshToken, auth.RefreshToken)
	if err != nil{
		return model.Tokens{}, err
	}
	if res.AccessToken, err = s.tokenManager.NewJWT(claims.UserID, claims.Role, s.accessTokenTTL, auth.AccessToken); err!= nil{
		return model.Tokens{}, err
	}
	return res, nil
}

func (s *UsersService) createSession(ctx context.Context, user *model.User) (model.Tokens, error){
	var (
		res model.Tokens
		err error
	)	
	res.AccessToken, err = s.tokenManager.NewJWT(user.ID, user.Role, s.accessTokenTTL, auth.AccessToken)
	if err != nil{
		return model.Tokens{}, err
	}
	res.RefreshToken, err = s.tokenManager.NewJWT(user.ID, user.Role, s.refreshTokenTTL, auth.RefreshToken)
	if err != nil{
		return model.Tokens{}, err
	}
	return res, nil
}

func (s *UsersService) GetUserIDFromToken(ctx context.Context, token string, tokenType string) (int, error){
	claims, err:= s.tokenManager.ParseToken(token, tokenType)
	if err != nil{
		return -1, err
	}
	return claims.UserID, nil
}

func (s *UsersService) Verify(accessToken string) error{
	return s.tokenManager.VerifyToken(accessToken, auth.AccessToken)
}