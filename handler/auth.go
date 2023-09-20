package handler

import (
	"Si-KP/database"
	"Si-KP/middleware"
	"Si-KP/response"
	"Si-KP/services"
	"Si-KP/table"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// @Summary Authentication
// @Description Authenticate user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body response.ReqAuth true "Authentication request"
// @Success 200 {object} response.GetUserLogin
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /login [post]
func Auth(c *fiber.Ctx) error {

	var request = new(response.ReqAuth)

	if err := c.BodyParser(&request); err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	// validasi auth pusri
	StatusCode := services.Auth_pusri_validation(request.Username, request.Password)
	// StatusCode := 200

	// kondisi success validasi
	if StatusCode == 200 {

		// save user login
		services.Save_user_login(request.Username)

		var user table.User_login
		database.DB.Where("username=?", request.Username).First(&user)

		token, err := middleware.GenerateJwt(strconv.Itoa(int(user.Id)))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return nil
		}

		// mengisi token ke kolom jwt
		isi_jwt := table.User_login{
			Jwt: token,
		}
		database.DB.Where("username=?", request.Username).Updates(&isi_jwt)

		c.Status(200)
		return c.JSON(fiber.Map{
			"message":      "Berhasil Login",
			"username":     request.Username,
			"role":         user.Role,
			"nama":         user.Name,
			"departemen":   user.Departemen,
			"access_token": token,
			"type_token":   "jwt",
			"auth_type":    "Bearer",
		})
	}

	// kondisi gagal validasi
	if StatusCode != 200 {

		var user table.Users
		var user_kp table.Users_Kp

		database.DB.Where("badge=?", request.Username).First(&user)
		// database.DB.Where("nim=?", request.Username).First(&user_kp)
		// database.DB.Unscoped().Where("badge=?", request.Username).First(&user)

		var password string
		if user.Id != 0 {
			password = user.Password
		} else {
			database.DB.Where("nim=?", request.Username).First(&user_kp)
			password = user_kp.Password
		}
		// password := user.Password // bcrypt password user
		hash := request.Password // inputan password
		match := services.CheckPasswordHash(password, hash)
		log.Println(match)
		log.Println(hash)
		log.Println(password)
		log.Println(user.Id)
		// if user.Deleted_by != "" {

		// }
		// if user.Badge != request.Username {
		// 	c.Status(400)
		// 	return c.JSON(fiber.Map{
		// 		"message": "Anda Belum Terdaftar",
		// 	})

		// }
		// if user.Deleted_by != "" {
		// 	c.Status(401)
		// 	return c.JSON(fiber.Map{
		// 		"message": "User has been deleted",
		// 	})
		// }
		if user.Id == 0 {
			database.DB.Where("nim= ?", request.Username).First(&user_kp)
			if user_kp.Id == 0 {
				c.Status(401)
				return c.JSON(fiber.Map{
					"message": "username atau password salah",
				})
			}

		}

		if !match {
			c.Status(401)
			return c.JSON(fiber.Map{
				"message": "password atau username salah",
			})
		}

		// kalau sukses
		c.Status(200)
		var user2 table.User_login
		database.DB.Where("username=?", request.Username).First(&user2)

		token, err := middleware.GenerateJwt(strconv.Itoa(int(user2.Id)))

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return nil
		}
		// mengisi token ke kolom jwt
		isi_jwt := table.User_login{
			Jwt: token,
		}
		database.DB.Where("username=?", request.Username).Updates(&isi_jwt)
		// save user login
		services.Save_user_login_local(request.Username)
		return c.JSON(fiber.Map{
			"message":      "Berhasil Login",
			"username":     request.Username,
			"role":         user2.Role,
			"nama":         user.Nama,
			"departemen":   user2.Departemen,
			"access_token": token,
			"type_token":   "jwt",
			"auth_type":    "Bearer",
		})
	}

	c.Status(401)
	return c.JSON(fiber.Map{
		"message": "Gagal Login",
	})

}

// Logout godoc
// @Summary Logout
// @Description Logout user
// @Tags Authentication
// @Security    ApiKeyAuth
// @Produce json
// @Success 200 {object} response.LogoutResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 405 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /logout [get]
func Logout(c *fiber.Ctx) error {

	db := database.DB
	token := c.Get("Authorization")

	if token == "" || (!strings.HasPrefix(token, "Bearer ") && token[7:] == "false") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized. Invalid or missing Bearer token.",
		})
	}
	if len(token) == 0 {
		// Jika token tidak ada, Anda dapat memberikan respons error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized. Token not found.",
		})
	}
	data := table.User_login{
		Jwt:          "-",
		Lates_logout: services.TimeNow(),
	}
	db.Where("username=?", services.GetUserLogin(c)).Updates(&data)

	return c.JSON(fiber.Map{
		"message": "Berhasil Logout",
	})

}
