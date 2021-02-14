package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xiaosasori/go-fiber/controllers"
)

// Setup setup
func Setup(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Post("/register", controllers.Register)
}
