package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func CreateApp() *fiber.App {

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100 mb limit uploud files
	})

	// Define the rate limiter middleware
	limiterConfig := limiter.Config{
		Max:        100,             // Maximum number of requests per duration
		Expiration: 1 * time.Minute, // Duration to keep records of requests
		KeyGenerator: func(c *fiber.Ctx) string {
			// Customize the rate limiter key if needed
			return c.IP() // Use client IP address as the key
		},
	}
	app.Use(limiter.New(limiterConfig))

	return app
}
