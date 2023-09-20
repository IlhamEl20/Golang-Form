package response

type GetUserLogin struct {
	Id         uint `json:"id"`
	Name       string
	Username   string
	Role       string
	Departemen string // unit kerja
	Dep_id     string // unit kerja id
	Email      string // email user login
	Jwt        string // json web token
}

type ReqAuth struct {
	Username string
	Password string
}
