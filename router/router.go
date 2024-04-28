package router

import (
	"app/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/v1", logger.New())
	api.Get("/", handler.Hello)

	// User
	user := api.Group("/user")
	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login)

}
