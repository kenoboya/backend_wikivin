package model

type BriefInfoProfile struct {
	Username  string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Image     string `json:"image" db:"image"`
}