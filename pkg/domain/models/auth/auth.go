package models

// Auth : model struct
type Auth struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username,required" db:"username"`
	Password string `json:"password" db:"password"`
	Fullname string `json:"fullname" db:"fullname"`
	Email    string `json:"email" db:"email"`
}
