package table

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id         uint `json:"id"`
	Role       string
	Created_at time.Time
	Created_by string
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Deleted_by string
}
