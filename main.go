package main

import (
	"Si-KP/config"
	"Si-KP/database"
	"Si-KP/routers"
	"Si-KP/services"
)

// @title Fiber Example API 2
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:9001
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	// database connect
	database.Connect()

	// drivex.WebdavConnect()

	// start fiber
	app := services.CreateApp()

	// router
	routers.Setup(app)

	// get port
	port := config.Env("PORT")

	// run
	app.Listen(":" + port)

}
