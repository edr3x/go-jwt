package initializers

import "edr3x/go-jwt/models"

func DbSync() {
	DB.AutoMigrate(&models.User{})
}
