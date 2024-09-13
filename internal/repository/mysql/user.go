package repo

import (
	"context"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository{
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, userSignUp model.UserSignUp) (model.User, error){
	var user model.User
	query:= "INSERT INTO users(username, email, password) VALUES(?,?,?)"
	
	result, err:= r.db.ExecContext(ctx, query, userSignUp.Username, userSignUp.Email, userSignUp.Password)
	if err != nil{
		return user, err
	}

	id, err:= result.LastInsertId()
	if err != nil{
		return user, err
	}
	if err:= r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = ?", id); err != nil{
		return model.User{}, err
	}
	return user, nil
}