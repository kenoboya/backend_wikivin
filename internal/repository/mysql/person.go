package repo

import (
	"context"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type PeopleRepository struct {
	db *sqlx.DB
}

func NewPeopleRepository(db *sqlx.DB) *PeopleRepository{
	return &PeopleRepository{db}
}

func (r *PeopleRepository) Create(ctx context.Context, person model.Person) error{
	query:= "INSERT INTO people(user_id, first_name, last_name, birth_date, gender, country, city, image) VALUES(:user_id, :first_name, :last_name, :birth_date, :gender, :country, :city, :image)"
	if _,err:= r.db.NamedExecContext(ctx, query, person); err!= nil{
		return err
	}
	return nil
}