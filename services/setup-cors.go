package services

import (
	"Si-KP/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: GetDomain(),
		// AllowOrigins:     "*",
		AllowHeaders:     config.AllowHeaders(),
		AllowMethods:     config.AllowMethods(),
		AllowCredentials: true,
	}))
}
