package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type router struct {
	*userHandler
	*catManagementHandler
	*catMatchHandler
}

func NewRouter(
	userHandler *userHandler,
	catMgtHandler *catManagementHandler,
	catMatchHandler *catMatchHandler,
) *router {
	return &router{
		userHandler: userHandler,
		catManagementHandler: catMgtHandler,
		catMatchHandler: catMatchHandler,
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

	cat := api.Group("/cat")
	cat.Post("", r.catManagementHandler.Create)
	cat.Get("", r.catManagementHandler.List)
	cat.Put("/:id", r.catManagementHandler.Update)
	cat.Delete("/:id", r.catManagementHandler.Delete)

	match := cat.Group("/match")
	match.Post("", r.catMatchHandler.Create)
}
