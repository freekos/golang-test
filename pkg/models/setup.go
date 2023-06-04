package models

import (
	"fmt"
	"freekos/crud/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		configs.Config.DB_HOST,
		configs.Config.DB_USER,
		configs.Config.DB_PASSWORD,
		configs.Config.DB_NAME,
		configs.Config.DB_PORT,
		configs.Config.DB_SSL_MODE,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	database.AutoMigrate(&Post{})

	DB = database
}
