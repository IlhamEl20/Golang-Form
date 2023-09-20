package table

import (
	"time"

	"gorm.io/gorm"
)

type Users_Kp struct {
	Id            uint   `json:"id"`
	Nim           string `json:"Nim"`
	Nama          string `json:"nama"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Status        string `json:"status"`
	Role_Id       int    `json:"role_id"`
	Periode_Id    int    `json:"periode_id"`
	// Status     string    `json:"status"`
	Penempatan string    `json:"penempatan"`
	Created_at time.Time `json:"created_at"`
	Created_by string    `json:"created_by"`
	Updated_at time.Time
	Updated_by string
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Deleted_by string
}
