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
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type catDetail struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	Description string    `json:"description"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageUrls   []string  `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

type listMatchesResponse struct {
	Id             string     `json:"id"`
	IssuedBy       userDetail `json:"issuedBy"`
	MatchCatDetail catDetail  `json:"matchCatDetail"`
	UserCatDetail  catDetail  `json:"userCatDetail"`
	Message        string     `json:"message"`
	CreatedAt      time.Time  `json:"createdAt"`
}

func (h *catMatchHandler) List(c *fiber.Ctx) error {
	matches, err := h.svc.List(c.Context(), "1")

	if err != nil {
		return serverError(c, fiber.StatusInternalServerError, "", err)
	}

	matchesResp := make([]listMatchesResponse, len(matches))
	for i, m := range matches {
		matchesResp[i] = listMatchesResponse{
			Id: m.Id,
			IssuedBy: userDetail{
				Name:      m.Name,
				Email:     m.Email,
				CreatedAt: m.UserCreatedAt,
			},
			MatchCatDetail: catDetail{
				Id: m.ReceiverCat.Id,
				Name: m.ReceiverCat.Name,
				Race: m.ReceiverCat.Race,
				Sex: m.ReceiverCat.Sex,
				Description: m.ReceiverCat.Description,
				AgeInMonth: m.ReceiverCat.AgeInMonth,
				ImageUrls: m.ReceiverCat.ImageUrls,
				HasMatched: m.ReceiverCat.HasMatched,
				CreatedAt: m.ReceiverCat.CreatedAt,
			},
			UserCatDetail: catDetail{
				Id: m.IssuerCat.Id,
				Name: m.IssuerCat.Name,
				Race: m.IssuerCat.Race,
				Sex: m.IssuerCat.Sex,
				Description: m.IssuerCat.Description,
				AgeInMonth: m.IssuerCat.AgeInMonth,
				ImageUrls: m.IssuerCat.ImageUrls,
				HasMatched: m.IssuerCat.HasMatched,
				CreatedAt: m.IssuerCat.CreatedAt,
			},
			Message: m.Message,
			CreatedAt: m.CatMatchCreatedAt,
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
