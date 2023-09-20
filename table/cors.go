package table

import (
	"time"

	"gorm.io/gorm"
)

type CorsDomain struct {
	Id         uint      `json:"id"`
	Domain     string    `json:"domain"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time
	Created_by string
	Updated_by string
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Deleted_by string
}
