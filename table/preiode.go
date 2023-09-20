package table

import (
	"time"

	"gorm.io/gorm"
)

type Periode struct {
	Id         uint           `json:"id"`
	Keterangan string         `json:"keterangan"`
	StartTime  time.Time      `gorm:"not null" json:"start_time"`
	EndTime    time.Time      `gorm:"not null" json:"end_time"`
	Created_at time.Time      `json:"created_at"`
	Created_by string         `json:"created_by"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
