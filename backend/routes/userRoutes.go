package routes

import (
	"godock/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUserByID) // Tambahkan ini
	app.Put("/users/:id", controllers.UpdateUser)  // Tambahkan ini
}
