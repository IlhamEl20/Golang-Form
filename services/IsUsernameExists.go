package services

import (
	"Si-KP/database"
)

func IsUsernameExists(badge string) (bool, error) {
	DB := database.DB
	var count int64

	result := DB.Table("users").Where("badge = ?", badge).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil

}
func IsUsernameExistsKp(nim string) (bool, error) {
	DB := database.DB
	var count int64

	result := DB.Table("users_kps").Where("nim = ?", nim).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil

}
