package api

import "github.com/gofiber/fiber/v2"

func failedToParseInput(c *fiber.Ctx, err error) error {
	return c.Status(500).JSON(fiber.Map{
		"status": "error",
		"message": "Failed to parse input",
		"errors": err.Error(),
	})
}

func invalidInput(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status": "error",
		"message": "Invalid input",
		"errors": err.Error(),
	})
}