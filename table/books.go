package table

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Nama       string    `json:"nama"`
	Title      string    `json:"title"`
	Judul      string    `json:"judul"`
	Penulis    string    `json:"penulis"`
	Penerbit   string    `json:"penerbit"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Created_by string
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
