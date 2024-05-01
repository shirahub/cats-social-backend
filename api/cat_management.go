package api

import (
	"app/port"
	"github.com/gofiber/fiber/v2"
)


type catManagementHandler struct {
	repo port.CatRepository
}

func NewCatManagementHandler(repo port.CatRepository) *catManagementHandler {
	return &catManagementHandler{repo}
}

func (h *catManagementHandler) Create(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
	})
}

func (h *catManagementHandler) List(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
	})
}

func (h *catManagementHandler) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
	})
}

func (h *catManagementHandler) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
	})
}