package api

import (
	"app/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/go-playground/validator/v10"
)

const iso8601 = "2006-01-02T15:04:05.999Z"

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

var validate = validator.New()

// SetupRoutes setup router api
func (r *router) Setup(app *fiber.App) {
	validate.RegisterValidation("oneof_races", validateRace)
	validate.RegisterValidation("compares_int", validateComparison)

	// Middleware
	api := app.Group("/v1", logger.New())

	// User
	user := api.Group("/user")
	user.Post("/register", r.userHandler.Register)
	user.Post("/login", r.userHandler.Login)

	cat := api.Group("/cat")
	cat.Post("", middleware.Protected(), r.catManagementHandler.Create)
	cat.Get("", middleware.Protected(), r.catManagementHandler.List)
	cat.Put("/:id", middleware.Protected(), r.catManagementHandler.Update)
	cat.Delete("/:id", middleware.Protected(), r.catManagementHandler.Delete)

	match := cat.Group("/match")
	match.Post("", middleware.Protected(), r.catMatchHandler.Create)
	match.Get("", middleware.Protected(), r.catMatchHandler.List)
	match.Post("/approve", middleware.Protected(), r.catMatchHandler.Approve)
	match.Post("/reject", middleware.Protected(), r.catMatchHandler.Reject)
	match.Delete("/:id", middleware.Protected(), r.catMatchHandler.Delete)
}
