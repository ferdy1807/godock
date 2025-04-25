package config

import (
	"fmt"
	"godock/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB menghubungkan aplikasi ke database PostgreSQL
func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=godock port=5432 sslmode=disable"
	var err error

	// Membuka koneksi ke database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Menampilkan status koneksi
	fmt.Println("Database connection successful")

	DB.AutoMigrate(&models.User{})
	return DB
}
