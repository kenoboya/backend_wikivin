package service

import (
	"context"
	"fmt"
	"time"
	"wikivin/internal/config"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
	"wikivin/pkg/auth"
	"wikivin/pkg/hash"
)

type Services struct {
	Articles Articles
	Users Users
	Profiles Profiles
	Favorites Favorites
}

func NewServices(deps Deps)*Services{
	return &Services{
		Articles: NewArticlesService(
			deps.repo.Articles, 
			deps.repo.Chapters, 
			deps.repo.InfoBox,
		),
		Users: NewUsersService(
			deps.repo.Users,
			deps.repo.People,
			deps.hasher,
			deps.tokenManager,
			deps.accessTokenTTL,
			deps.refreshTokenTTL,
		),
		Profiles: NewProfileService(
			deps.repo.Profiles,
		),
		Favorites: NewFavoritesService(
			deps.repo.Favorites,
		),
	}
}

type Deps struct{
	repo *repo.Repositories
	hasher hash.PasswordHasher
	tokenManager auth.TokenManager
	accessTokenTTL time.Duration
	refreshTokenTTL time.Duration
}

func NewDeps(repo *repo.Repositories, config config.AuthConfig)(*Deps, error){
	hasher:= hash.NewSHA256Hasher(config.PasswordSalt)
	tokenManager, err:= auth.NewManager(config.JWT.SecretAccessKey, config.JWT.SecretRefreshKey)
	if err != nil{
		return nil, fmt.Errorf("tokenManager: %w", err)
	}
	return &Deps{
			repo: repo,
			hasher: hasher,
			tokenManager: tokenManager,
			accessTokenTTL: config.JWT.AccessTokenTTL,
			refreshTokenTTL: config.JWT.RefreshTokenTTL,
		},nil
	}

type Articles interface {
	CreateArticle(ctx context.Context, infoBoxDB model.InfoBoxDB, article model.Article, chapters []model.Chapter) error
	LoadArticles(ctx context.Context) ([]model.Article, error)
	LoadArticle(ctx context.Context, title string) (*model.ArticlePage, error)
}

type Users interface{
	SignUp(ctx context.Context, requestSignUp model.RequestSignUp)(model.Tokens, error)
	SignIn(ctx context.Context, requestSignIn model.UserSignIn)(model.Tokens, error)
	RefreshToken(ctx context.Context, refreshToken string)(model.Tokens, error)
	
	GetUserIDFromToken(ctx context.Context, token string, tokenType string) (int, error)
	GetAccessTokenTTL() time.Duration
	GetRefreshTokenTTL() time.Duration
}

type Profiles interface{
	GetBriefInfoProfile(ctx context.Context, userID int) (model.BriefInfoProfile, error)
}

type Favorites interface{
	GetFavoriteArticlesByUserID(ctx context.Context, userID int)([]model.FavoriteArticle, error)
}
