package table

import (
	"time"

	"gorm.io/gorm"
)

type Silabus struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Nama       uint      `json:"nama"`
	Title      string    `json:"title"`
	Jenis      string    `json:"jenis"`
	Tahun      time.Time `json:"tahun"`
	Start_date time.Time `json:"start_date"`
	End_date   time.Time `json:"end_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Created_by string
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
