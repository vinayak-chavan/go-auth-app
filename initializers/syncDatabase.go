package initializers

import "go-auth-app/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}