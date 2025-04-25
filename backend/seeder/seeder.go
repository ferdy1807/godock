package seeder

import (
	"godock/config"
	"log"
	"time"
)

// User struct defines the model for users table
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100"`
	Birthdate time.Time `gorm:"type:date"`
	Gender    string    `gorm:"size:10"`
	Job       string    `gorm:"size:100"`
	Photo     string    `gorm:"size:255"`
}

// ResetDatabase akan menghapus semua data dalam tabel users
// func ResetDatabase() {
// 	// Menghapus seluruh data di tabel users
// 	if err := config.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error; err != nil {
// 		log.Fatal("Failed to reset database:", err)
// 	} else {
// 		log.Println("Database successfully reset")
// 	}
// }

// CreateTableIfNotExists will auto-migrate the schema for users table
func CreateTableIfNotExists() {
	// GORM AutoMigrate will create the table if it doesn't exist
	if err := config.DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("Failed to create table:", err)
	} else {
		log.Println("Table 'users' created or already exists")
	}
}

// SeedUsers will insert predefined users into the 'users' table
func SeedUsers() {
	users := []User{
		{Name: "Ali", Birthdate: time.Date(1990, 1, 15, 0, 0, 0, 0, time.UTC), Gender: "Laki-laki", Job: "Programmer", Photo: "ali.jpg"},
		{Name: "Budi", Birthdate: time.Date(1995, 6, 10, 0, 0, 0, 0, time.UTC), Gender: "Laki-laki", Job: "Designer", Photo: "budi.jpg"},
		{Name: "Citra", Birthdate: time.Date(1992, 3, 22, 0, 0, 0, 0, time.UTC), Gender: "Perempuan", Job: "Manager", Photo: "citra.jpg"},
	}

	// Iterate over the users and insert them using GORM's Create method
	for _, user := range users {
		if err := config.DB.Create(&user).Error; err != nil {
			log.Printf("Failed to insert user %s: %v", user.Name, err)
		} else {
			log.Printf("Inserted user %s", user.Name)
		}
	}
}
