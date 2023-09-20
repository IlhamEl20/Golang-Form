package services

import (
	"Si-KP/database"
	"Si-KP/table"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRoleLogin(c *fiber.Ctx) string {

	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var user table.User_login
	database.DB.Where("jwt = ?", token).First(&user)

	return user.Role

}
