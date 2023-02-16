package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(
		postgres.Open(os.Getenv("DB")),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate()
	DB = db
}
