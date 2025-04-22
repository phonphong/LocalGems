package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:15022002@tcp(127.0.0.1:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Hoặc dùng SQLite:
	// database, err := gorm.Open(sqlite.Open("places.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Không kết nối được DB: ", err)
	}

	DB = database
}