package api

import "github.com/gofiber/fiber/v2"

func serverError(c *fiber.Ctx, statusCode int, message string, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status": "error",
		"message": message,
		"errors": err.Error(),
	})
}

func failedToParseInput(c *fiber.Ctx, err error) error {
	return serverError(c, fiber.StatusInternalServerError, "Failed to parse input", err)
}

func invalidInput(c *fiber.Ctx, err error) error {
	return serverError(c, fiber.StatusBadRequest, "Invalid input", err)
}