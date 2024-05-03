package api

import (
	"app/domain"
	"errors"
	"github.com/gofiber/fiber/v2"
)

var errorCodes = map[error]int{
	domain.ErrNotFound: fiber.StatusNotFound,
	domain.ErrEmailTaken: fiber.StatusConflict,
	domain.ErrCatInMatch: fiber.StatusBadRequest,
	domain.ErrMatchResponded: fiber.StatusBadRequest,
	domain.ErrMatchWithOwnedCat: fiber.StatusBadRequest,
	domain.ErrMatchWithSameSex: fiber.StatusBadRequest,
	domain.ErrMatchWithTaken: fiber.StatusBadRequest,
	domain.ErrMatchExists: fiber.StatusBadRequest,
}

func serverError(c *fiber.Ctx, err error) error {
	for key, status := range errorCodes {
		if errors.Is(err, key) {
			return customError(c, status, "", err)
		}
	}

	return customError(c, fiber.StatusInternalServerError, "", err)
}

func customError(c *fiber.Ctx, statusCode int, message string, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status": "error",
		"message": message,
		"errors": err.Error(),
	})
}

func failedToParseInput(c *fiber.Ctx, err error) error {
	return customError(c, fiber.StatusBadRequest, "Failed to parse input", err)
}

func invalidInput(c *fiber.Ctx, err error) error {
	return customError(c, fiber.StatusBadRequest, "Invalid input", err)
}