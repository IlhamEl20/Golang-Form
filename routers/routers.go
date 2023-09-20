package routers

import (
	"Si-KP/config"
	"Si-KP/docs"
	"Si-KP/handler"
	"Si-KP/middleware"
	"Si-KP/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Setup(app *fiber.App) {

	// setup cors
	services.SetupCors(app)

	app.Get("/docs/*", swagger.HandlerDefault) // default
	docs.SwaggerInfo.Host = config.Env("HOST")
	app.Get("/", handler.AppName)
	app.Post("/login", handler.Auth)

	// route dibawah ini wajib login
	// app.Use(middleware.IsAuthenticate)
	app.Get("/logout", handler.Logout)

	// users
	app.Post("/users", middleware.IsAuthenticate("ada", "Admin"), handler.CreateUsers)
	app.Put("/users/:id", middleware.IsAuthenticate("ada", "Admin"), handler.UpdateUser)
	app.Get("/users", middleware.IsAuthenticate("ada", "Admin"), handler.GetUsers)
	app.Delete("/users/:id", middleware.IsAuthenticate("ada", "Admin"), handler.SoftDeleteUser)

	//Role
	// app.Post("/role-create", middleware.IsAuthenticate("Admin"), handler.CreateRole)
	app.Get("/role", middleware.IsAuthenticate("Admin"), handler.GetRole)
	// app.Put("/role/:id", middleware.IsAuthenticate("Admin"), handler.UpdateRole)
	// app.Delete("/role/:id", middleware.IsAuthenticate("Admin"), handler.SoftDeleteRole)

	// Books
	app.Post("/books", middleware.IsAuthenticate("Admin"), handler.CreateBooks)
	app.Get("/books", middleware.IsAuthenticate("Admin"), handler.GetBooks)
	app.Put("/books/:id", middleware.IsAuthenticate("Admin"), handler.UpdateBooks)
	app.Delete("/books/:id", middleware.IsAuthenticate("Admin"), handler.SoftDeleteBooks)

}
