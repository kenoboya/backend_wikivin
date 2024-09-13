package service

import (
	"context"
	"wikivin/internal/model"
	repo "wikivin/internal/repository/mysql"
)

type Services struct {
	Articles Articles
	User User
}

func NewServices(repo *repo.Repositories)*Services{
	return &Services{
		Articles: NewArticlesService(repo.Articles, repo.Chapters, repo.InfoBox),
		User: NewUsersService(
			repo.Users,
			repo.People,
			hasher,
			tokenManager,
		),
	}
}

type Articles interface {
	CreateArticle(ctx context.Context, infoBoxDB model.InfoBoxDB, article model.Article, chapters []model.Chapter) error
	LoadArticles(ctx context.Context) ([]model.Article, error)
	LoadArticle(ctx context.Context, title string) (*model.ArticlePage, error)
}

type User interface{
	SignUp(ctx context.Context, requestSignUp model.RequestSignUp)(model.Tokens, error)
	SignIn(ctx context.Context, requestSignIn model.UserSignIn)(model.Tokens, error)
	RefreshToken(ctx context.Context, refreshToken string)(model.Tokens, error)
}
