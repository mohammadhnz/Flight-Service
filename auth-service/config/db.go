package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(
		postgres.Open("postgres://postgres:postgres@localhost:8081/postgres"),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate()
	DB = db
}
