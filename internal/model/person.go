package model

const (
	Male   = "male"
	Female = "female"
	Other  = "other"
)

type Person struct {
	ID        int     `json:"person_id" db:"person_id"`
	UserID    int     `json:"user_id,omitempty" db:"user_id"`
	FirstName string  `json:"first_name" db:"first_name"`
	LastName  string  `json:"last_name" db:"last_name"`
	BirthDay  string  `json:"birth_day" db:"birth_day"`
	Gender    string  `json:"gender" db:"gender"`
	County    *string `json:"country" db:"country"`
	City      *string `json:"city" db:"city"`
	Image     *string `json:"image" db:"image"`
}