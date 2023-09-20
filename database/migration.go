package database

import (
	"Si-KP/table"
)

func Migration() {

	DB.AutoMigrate(

		&table.Users{},
		&table.Role{},
		&table.User_login{},
		&table.CorsDomain{},
		&table.Periode{},
		&table.Lampiran_Kp{},
		&table.Users_Kp{},
		&table.Books{},
		&table.Atk{},
		&table.Silabus{},
	)

}
