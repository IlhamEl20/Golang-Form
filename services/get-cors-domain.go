package services

import (
	"Si-KP/database"
	constants "Si-KP/global-variable"
	"Si-KP/table"
	"log"
)

func GetDomain() string {

	db := database.DB

	var (
		data   []table.CorsDomain
		result string
	)

	if err := db.Table(constants.TABLE_CORS_DOMAIN).Find(&data).Error; err != nil {
		log.Println("Failed to retrieve data from the database:", err)
	}

	for i, domain := range data {
		if i > 0 {
			result += ", "
		}
		result += domain.Domain
	}

	return result

}
