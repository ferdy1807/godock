package controllers

import (
	"godock/models"
	"log"
	"strconv" // Untuk konversi string ke integer

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init digunakan untuk menginisialisasi database
func Init(database *gorm.DB) {
	db = database
}

// GetUsers mengambil semua user dari database
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	// Menggunakan db untuk query (menghindari penggunaan config.DB secara langsung)
	if err := db.Find(&users).Error; err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}
	return c.JSON(users)
}

// GetUserByID mengambil user berdasarkan ID dari database
func GetUserByID(c *fiber.Ctx) error {
	// Ambil ID dari URL parameter, konversi ke integer
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr) // Mengonversi string ke integer
	if err != nil {
		// Jika konversi gagal, kembalikan error
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var user models.User
	// Gunakan GORM untuk mencari user berdasarkan ID
	if err := db.First(&user, id).Error; err != nil {
		// Jika user tidak ditemukan, kembalikan status 404
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// UpdateUser memperbarui data user berdasarkan ID
func UpdateUser(c *fiber.Ctx) error {
	// Ambil ID dari URL parameter
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr) // Mengonversi string ke integer
	if err != nil {
		// Jika konversi gagal, kembalikan error
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var user models.User
	// Cari user berdasarkan ID
	if err := db.First(&user, id).Error; err != nil {
		// Jika user tidak ditemukan, kembalikan status 404
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Parsing data dari request body untuk update
	if err := c.BodyParser(&user); err != nil {
		// Jika data invalid, kembalikan status 400
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	// Simpan perubahan user ke database
	if err := db.Save(&user).Error; err != nil {
		// Jika terjadi error saat menyimpan, kembalikan status 500
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(user)
}
