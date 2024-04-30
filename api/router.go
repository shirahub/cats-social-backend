package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type router struct {
	*userHandler
}

func NewRouter(userHandler *userHandler) *router {
	return &router{
		userHandler: userHandler,
	}
}

// SetupRoutes setup router api
func (r *router) Setup(app *fiber.App) {
	// Middleware
	api := app.Group("/v1", logger.New())

	// User
	user := api.Group("/user")
	user.Post("/register", r.userHandler.Register)
	user.Post("/login", r.userHandler.Login)

}
