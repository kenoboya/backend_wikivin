package model

const (
	Blocked   = "blocked"
	Unblocked = "unblocked"

	Active   = "active"
	Inactive = "inactive"

	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	ID           int    `json:"user_id" db:"user_id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	Password     string `json:"password" db:"password"`
	Status       string `json:"status" db:"status"`
	Blocked      string `json:"blocked" db:"blocked"`
	RegisteredAt string `json:"registered_at" db:"registered_at"`
	LastLogin    string `json:"last_login" db:"last_login"`
	Role         string `json:"role" db:"role"`
}

func (u *User) IsBlocked() bool {
	return u.Blocked == Blocked
}

type UserSignUp struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
type UserSignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RequestSignUp struct {
	UserSignUp UserSignUp `json:"UserSignUp"`
	Person     Person     `json:"Person"`
}