package configs

import (
	"20_Clean_Architecture/praktikum/models"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}