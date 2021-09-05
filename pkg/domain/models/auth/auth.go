package models

// Auth : model struct
type Auth struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" validate:"required" db:"username"`
	Password string `json:"password" validate:"required" db:"password"`
	Fullname string `json:"fullname" validate:"required" db:"fullname"`
	Email    string `json:"email" validate:"required" db:"email"`
}
