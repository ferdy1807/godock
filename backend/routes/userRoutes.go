package routes

import (
	"godock/backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
}
