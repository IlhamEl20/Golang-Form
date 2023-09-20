package table

import "time"

type User_login struct {
	Id           uint `json:"id"`
	Name         string
	Username     string
	Role         string
	Departemen   string // unit kerja
	Email        string // email user login
	Jwt          string // json web token
	Lates_login  time.Time
	Lates_logout time.Time
}
