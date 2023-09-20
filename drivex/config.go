package drivex

import (
	"Si-KP/config"
	"log"

	"github.com/studio-b12/gowebdav"
)

func WebdavConnect() *gowebdav.Client {
	host := config.Env("DRIVEX_HOST")
	user := config.Env("DRIVEX_USERNAME")
	password := config.Env("DRIVEX_PASS")
	c := gowebdav.NewClient(host, user, password)
	// return c
	// c.Connect()
	if c.Connect() != nil {
		panic("Failed connect to webdev")
	} else {
		log.Println("Connection to webdev success!")
	}
	return c
}
