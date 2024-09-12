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

func (r *UserRepository) Create(ctx context.Context, user model.UserSignUp) (int, error){
	query:= "INSERT INTO users(username, email, password) VALUES(?,?,?)"
	
	result, err:= r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil{
		return -1, err
	}
	id, err:= result.LastInsertId()
	if err != nil{
		return -1, err
	}
	return int(id), nil
}