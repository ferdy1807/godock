package routes

import (
	"godock/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/count", controllers.Home)
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUserByID) // Menangani rute untuk mendapatkan user berdasarkan ID
	app.Put("/users/:id", controllers.UpdateUser)  // Menangani rute untuk memperbarui user berdasarkan ID
}
