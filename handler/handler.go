package handler

import (
	// constants "Si-KP/global-variable"

	"github.com/gofiber/fiber/v2"
)

func AppName(c *fiber.Ctx) error {
	c.Status(200)
	return c.Redirect("/docs/index.html", fiber.StatusFound)
	// return c.JSON(fiber.Map{

	// 	// "app_name": constants.APP_NAME,
	// 	// "desc":     constants.DESC,
	// })

}
