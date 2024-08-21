package database

import (
	"github.com/sanchit-sherawat/go-backend.git/configs"
	"github.com/sanchit-sherawat/go-backend.git/models"
)

func Migrate() {
	configs.DB.AutoMigrate(&models.User{})
}
