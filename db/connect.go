package db

import (
	"dev-diary/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func Connect() error {
	db, err := gorm.Open(sqlite.Open("devdiary.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	DB.AutoMigrate(models.Post{})
	return nil
}
