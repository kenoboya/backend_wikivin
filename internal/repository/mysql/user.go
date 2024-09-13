package repo

import (
	"context"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository{
	return &UsersRepository{db}
}

func (r *UsersRepository) Create(ctx context.Context, userSignUp model.UserSignUp) (model.User, error){
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

func (r *UsersRepository) GetByLogin(ctx context.Context, login string) (model.User, error){
	var user model.User
	query:= "SELECT * FROM users WHERE email = ? OR username = ?"
	if err:= r.db.GetContext(ctx, &user, query, login, login); err != nil{
		return model.User{}, err
	}
	return user, nil
}