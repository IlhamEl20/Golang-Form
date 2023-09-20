package response

import (
	"time"
)

type GetUsers struct {
	Id         uint      `json:"id"`
	Badge      string    `json:"badge"`
	Nama       string    `json:"nama"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Created_by string    `json:"created_by"`
	Role_Id    int       `json:"role_id"`
	Role       string    `json:"role"`
}

type FormUsers struct {
	Badge    string `json:"badge" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	Role_Id  int    `json:"role_id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type LogoutResponse struct {
	Message string `json:"message"`
}
