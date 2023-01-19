package models

type Users struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
