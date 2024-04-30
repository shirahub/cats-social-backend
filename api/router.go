package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type router struct {
	*userHandler
	*catManagementHandler
}

func NewRouter(userHandler *userHandler, catMgtHandler *catManagementHandler) *router {
	return &router{
		userHandler: userHandler,
		catManagementHandler: catMgtHandler,
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

	catManagement := api.Group("/cat")
	catManagement.Post("", r.catManagementHandler.Create)
	catManagement.Get("", r.catManagementHandler.List)
	catManagement.Put("/:id", r.catManagementHandler.Update)
	catManagement.Delete("/:id", r.catManagementHandler.Delete)
}
