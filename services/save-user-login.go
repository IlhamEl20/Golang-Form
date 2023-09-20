package services

import (
	"Si-KP/config"
	"Si-KP/database"
	"Si-KP/table"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// simpan data login user iam pusri
func Save_user_login(username string) error {

	url_master_karyawan := config.Env("URL_MASTER_KARYAWAN")

	var (
		role_login string
		role_user  table.Users
		role       table.Role
	)

	database.DB.Where("badge=?", username).First(&role_user)

	var cek_user int64
	database.DB.Model(&table.Users{}).Where("badge = ?", username).Count(&cek_user)

	if cek_user > 0 {
		database.DB.Where("id=?", role_user.Role_Id).First(&role)
		role_login = role.Role
	}

	if cek_user < 1 {
		database.DB.Where("id=?", 1).First(&role)
		role_login = role.Role
	}

	resp, err := http.Get(url_master_karyawan + "&emp_no=6" + username)
	if err != nil {
		log.Println(err)
	}

	var res map[string][]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	nama := ""
	email := ""
	departemen := ""

	for _, v := range res["data"] {

		nama = v.(map[string]interface{})["nama"].(string)
		email = v.(map[string]interface{})["email"].(string)
		departemen = v.(map[string]interface{})["dept_title"].(string)

	}

	user := table.User_login{
		Name:        nama,
		Username:    username,
		Email:       email,
		Departemen:  departemen,
		Role:        role_login,
		Lates_login: time.Now().Local(),
	}

	var total int64
	database.DB.Model(&table.User_login{}).Where("username = ?", username).Count(&total)

	if total < 1 {

		err := database.DB.Create(&user)

		if err != nil {
			log.Println(err)
		}
	}

	if total > 0 {
		database.DB.Where("username = ?", username).Updates(&user)
	}

	return nil
}

func Save_user_login_local(username string) error {

	var (
		user_detail table.Users
		role_login  string
		role        table.Role
	)

	database.DB.Where("badge = ?", username).First(&user_detail)

	database.DB.Where("id=?", user_detail.Role_Id).First(&role)

	role_login = role.Role

	user := table.User_login{

		Name:        user_detail.Nama,
		Username:    user_detail.Badge,
		Email:       user_detail.Email,
		Role:        role_login,
		Departemen:  "",
		Lates_login: time.Now().Local(),
	}

	var total int64
	database.DB.Model(&table.User_login{}).Where("username = ?", username).Count(&total)

	if total < 1 {

		err := database.DB.Create(&user)

		if err != nil {
			log.Println(err)
		}
	}

	if total > 0 {
		database.DB.Where("username = ?", username).Updates(&user)
	}

	return nil
}
