package api

import (
	"app/domain"
	"app/port"
	"github.com/gofiber/fiber/v2"
)

type catMatchHandler struct {
	svc port.CatMatchService
}

func NewCatMatchHandler(svc port.CatMatchService) *catMatchHandler {
	return &catMatchHandler{svc}
}

type createMatchRequest struct {
	MatchCatId string
	UserCatId  string
	Message    string `validate:"min=5,max=120"`
}

func (h *catMatchHandler) Create(c *fiber.Ctx) error {
	req := new(createMatchRequest)
	if err := c.BodyParser(req); err != nil {
		return failedToParseInput(c, err)
	}

	if err := validate.Struct(req); err != nil {
		return invalidInput(c, err)
	}

	match := domain.CatMatch{
		IssuerCatId: req.UserCatId,
		ReceiverCatId: req.MatchCatId,
		Message: req.Message,
	}

	newRecord, err := h.svc.Create(c.Context(), &match)
	if err != nil {
		return serverError(c, fiber.StatusInternalServerError, "", err)
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"id": newRecord.Id,
			"createdAt": newRecord.CreatedAt.Format(iso8601),
		},
	})
}

func (h *catMatchHandler) List(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

type updateMatchRequest struct {
	MatchId string
}

func (h *catMatchHandler) Approve(c *fiber.Ctx) error {
	req := new(updateMatchRequest)
	if err := c.BodyParser(req); err != nil {
		return failedToParseInput(c, err)
	}

	if err := validate.Struct(req); err != nil {
		return invalidInput(c, err)
	}

	matchId, updatedAt, err := h.svc.Approve(req.MatchId, "1")
	if err != nil {
		if err == domain.ErrNotFound {
			return serverError(c, fiber.StatusNotFound, "", err)
		}
		if err == domain.ErrMatchResponded {
			return serverError(c, fiber.StatusBadRequest, "", err)
		}
		return serverError(c, fiber.StatusInternalServerError, "", err)
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"id":        matchId,
			"updatedAt": updatedAt.Format(iso8601),
		},
	})
}

func (h *catMatchHandler) Reject(c *fiber.Ctx) error {
	req := new(updateMatchRequest)
	if err := c.BodyParser(req); err != nil {
		return failedToParseInput(c, err)
	}

	if err := validate.Struct(req); err != nil {
		return invalidInput(c, err)
	}

	matchId, updatedAt, err := h.svc.Reject(req.MatchId, "1")
	if err != nil {
		if err == domain.ErrNotFound {
			return serverError(c, fiber.StatusNotFound, "", err)
		}
		if err == domain.ErrMatchResponded {
			return serverError(c, fiber.StatusBadRequest, "", err)
		}
		return serverError(c, fiber.StatusInternalServerError, "", err)
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"id":        matchId,
			"updatedAt": updatedAt.Format(iso8601),
		},
	})
}

func (h *catMatchHandler) Delete(c *fiber.Ctx) error {
	id, deletedAt, err := h.svc.Delete(c.Context(), c.Params("id"), "1")
	if err != nil {
		if err == domain.ErrNotFound {
			return serverError(c, fiber.StatusNotFound, "", err)
		}
		if err == domain.ErrMatchResponded {
			return serverError(c, fiber.StatusBadRequest, "", err)
		}
		return serverError(c, fiber.StatusInternalServerError, "", err)
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"id": id,
			"deletedAt": deletedAt,
		},
	})
}
