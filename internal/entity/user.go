package entity

type User struct {
	ID       int64  `json:"id,omitempty" db:"id"`
	Username string `json:"username,omitempty" db:"username"`
	Email    string `json:"email" db:"email"`
	Status   int    `json:"status,omitempty" db:"status"`
	Role     string `json:"role" db:"role"`
}
