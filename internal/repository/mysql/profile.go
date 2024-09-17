package repo

import (
	"context"
	"fmt"
	"wikivin/internal/model"

	"github.com/jmoiron/sqlx"
)

type ProfileRepository struct {
	db *sqlx.DB
}

func NewProfileRepository(db *sqlx.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (r *ProfileRepository) GetBriefInfoProfile(ctx context.Context, userID int) (model.BriefInfoProfile, error) {
	var briefInfoProfile model.BriefInfoProfile
	query := `
		SELECT u.username, p.first_name, p.last_name, p.image
		FROM people p
		INNER JOIN users u ON p.user_id = u.user_id
		WHERE p.user_id = ?`
	err := r.db.Get(&briefInfoProfile, query, userID)
	if err != nil {
		fmt.Println(err)
		return model.BriefInfoProfile{}, err
	}
	return briefInfoProfile, nil
}
