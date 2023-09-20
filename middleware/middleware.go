package middleware

import (
	"Si-KP/database"
	"Si-KP/table"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticate(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token auth
		token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

		var user table.User_login
		database.DB.Where("jwt = ?", token).First(&user)
		// log.Println(user.Jwt)

		// claim token, aktif atau kadaluarsa
		if _, err := Parsejwt(user.Jwt); err != nil {

			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "unauthenticated || Access Denied",
			})
		}

		// cek ke database login
		if user.Jwt == "" || user.Jwt != token {

			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "unauthenticated || Access Denied",
			})
		}

		isRoleAllowed := false
		for _, allowedRole := range allowedRoles {
			if user.Role == allowedRole {
				isRoleAllowed = true
				// log.Println(allowedRoles)

				break

			}
		}

		if !isRoleAllowed {
			// Return a 403 Forbidden status if the user's role is not allowed
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}
		return c.Next()
	}
}
