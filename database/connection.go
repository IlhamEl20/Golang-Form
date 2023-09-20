package database

import (
	"Si-KP/config"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	host := config.Env("DB_HOST")
	port := config.Env("DB_PORT")
	user := config.Env("DB_USERNAME")
	password := config.Env("DB_PASSWORD")
	name := config.Env("DB_NAME")

	// Connect to data-express Database
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, name)

	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect to database")
	} else {
		log.Println("Connection to Database success!")
	}

	Migration()
}
