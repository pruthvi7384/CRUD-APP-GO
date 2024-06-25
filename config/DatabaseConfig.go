package config

import (
	"crudApp/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DataBase Connection

var DB *gorm.DB

func DatabaseConfg() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"
	databae, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[DatabaseConfig] Failed To Connect Database!")
		panic(err)
	} else {
		log.Printf("[DatabaseConfig] Database Conection Done Successfully!")
	}

	// Database Initilization
	DB = databae

	// AutoMigration Set Here
	DB.AutoMigrate(&models.User{})
}
