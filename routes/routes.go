package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xiaosasori/go-fiber/controllers"
	"github.com/xiaosasori/go-fiber/middlewares"
)

// Setup setup
func Setup(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", controllers.Home)
	// Auth
	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	// User
	user := api.Group("/user")
	user.Get("/:id", controllers.GetUser)
	user.Patch("/", middlewares.Protected(), controllers.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), handler.DeleteUser)
}
