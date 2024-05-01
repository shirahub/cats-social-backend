package api

import (
	"app/port"
	"github.com/gofiber/fiber/v2"
)


type catMatchHandler struct {
	repo port.MatchRepository
}

func NewCatMatchHandler(repo port.MatchRepository) *catMatchHandler {
	return &catMatchHandler{repo}
}

func (h *catMatchHandler) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
	})
}
