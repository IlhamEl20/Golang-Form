package table

import (
	"time"

	"gorm.io/gorm"
)

type Lampiran_Kp struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Nim        uint      `json:"nim" gorm:"not null;unique"`
	Slug       string    `json:"slug"`
	Nama_File  string    `json:"nama_file" gorm:"type:varchar(100);not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Created_by string
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
