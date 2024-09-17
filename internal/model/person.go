package model

const (
	Male   = "male"
	Female = "female"
	Other  = "other"
)

type Person struct {
	ID        int     `json:"person_id,omitempty" db:"person_id"`
	UserID    int     `json:"user_id,omitempty" db:"user_id"`
	FirstName string  `json:"first_name" db:"first_name"`
	LastName  string  `json:"last_name" db:"last_name"`
	BirthDate string  `json:"birth_date" db:"birth_date"`
	Gender    string  `json:"gender" db:"gender"`
	County    *string `json:"country" db:"country"`
	City      *string `json:"city" db:"city"`
	Image     *string `json:"image" db:"image"`
}