package main

import (
	"godock/config"
	"godock/controllers"
	"godock/routes"
	"godock/seeder"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Ambil koneksi DB dari config
	db := config.ConnectDB()

	// Inisialisasi controller dengan DB
	controllers.Init(db)

	// Seed data (opsional)
	seeder.CreateTableIfNotExists()
	// seeder.ResetDatabase()
	seeder.SeedUsers()

	// Setup routes
	routes.UserRoutes(app)

	// Start server
	app.Listen(":8080")
}
