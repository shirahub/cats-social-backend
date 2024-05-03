package api

import (
	"app/domain"
	"app/port"
	"time"
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

type userDetail struct {
	Name string
	Email string
	CreatedAt time.Time
}

type catDetail struct {
	Id string
	Name string
	Race string
	Sex string
	Description string
	AgeInMonth int
	ImageUrls []string
	HasMatched bool
	CreatedAt time.Time
}

type listMatchesResponse struct {
	Id string
	IssuedBy userDetail
	MatchCatDetail catDetail
	UserCatDetail catDetail
	Message string
	CreatedAt time.Time
}

func (h *catMatchHandler) List(c *fiber.Ctx) error {
	matches, err := h.svc.List(c.Context())

	if err != nil {
		return serverError(c, fiber.StatusInternalServerError, "", err)
	}

	matchesResp := make([]listMatchesResponse, len(matches))
	for i, m := range matches {
		matchesResp[i] = listMatchesResponse{
			Id: m.Id,
			IssuedBy: userDetail{
				Email: m.Email,
			},
			MatchCatDetail: catDetail{
				Name: m.ReceiverCat.Name,
				Race: m.ReceiverCat.Race,
				Sex: m.ReceiverCat.Sex,
				Description: m.ReceiverCat.Description,
				AgeInMonth: m.ReceiverCat.AgeInMonth,
				ImageUrls: m.ReceiverCat.ImageUrls,
				HasMatched: m.ReceiverCat.HasMatched,
				CreatedAt: m.ReceiverCat.CreatedAt,
			},
		}
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data": matchesResp,
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

	matchId, updatedAt, err := h.svc.Approve(c.Context(), req.MatchId, "1")
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

	matchId, updatedAt, err := h.svc.Reject(c.Context(), req.MatchId, "1")
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
