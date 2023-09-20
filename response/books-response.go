package response

import (
	"time"
)

type GetBooks struct {
	Id         uint      `json:"id"`
	Nama       string    `json:"nama"`
	Title      string    `json:"title"`
	Judul      string    `json:"judul"`
	Penulis    string    `json:"penulis"`
	Penerbit   string    `json:"penerbit"`
	Created_at time.Time `json:"created_at"`
	Created_by string    `json:"created_by"`
}
type FormBooks struct {
	Nama     string `json:"nama" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Judul    string `json:"judul" validate:"required"`
	Penulis  string `json:"penulis" validate:"required"`
	Penerbit string `json:"penerbit" validate:"required"`
}
