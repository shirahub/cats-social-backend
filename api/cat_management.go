package api

import (
	"app/domain"
	"app/port"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type catManagementHandler struct {
	svc port.CatManagementService
}

func NewCatManagementHandler(svc port.CatManagementService) *catManagementHandler {
	return &catManagementHandler{svc}
}

type createUpdateCatRequest struct {
	Name        string   `validate:"min=1,max=30"`
	Race        string   `validate:"oneof_races"`
	Sex         string   `validate:"oneof=male female"`
	AgeInMonth  int      `validate:"min=1,max=120082"`
	Description string   `validate:"min=1,max=200"`
	ImageUrls   []string `validate:"min=1,max=10"`
}

type listCatsRequest struct {
	Id         string
	Limit      int    `validate:"gt=0"`
	Offset     int
	Race       string
	Sex        string `validate:"omitempty,oneof=male female"`
	HasMatched string `validate:"omitempty,boolean"`
	AgeInMonth string `validate:"omitempty,compares_int"`
	Owned      string `validate:"omitempty,boolean"`
	Name       string
}

func (h *catManagementHandler) Create(c *fiber.Ctx) error {
	req := new(createUpdateCatRequest)
	if err := c.BodyParser(req); err != nil {
		return failedToParseInput(c, err)
	}

	if err := validate.Struct(req); err != nil {
		return invalidInput(c, err)
	}

	cat := domain.CreateCatRequest{
		Name:        req.Name,
		Race: 			 req.Race,
		Sex:         req.Sex,
		AgeInMonth:  req.AgeInMonth,
		Description: req.Description,
		ImageUrls:   req.ImageUrls,
		UserId:      "1",
	}
	newRecord, err := h.svc.Create(&cat)
	if err != nil {
		return serverError(c, fiber.StatusInternalServerError, "", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"id":        newRecord.Id,
			"createdAt": newRecord.CreatedAt.Format(iso8601),
		},
	})
}

func (h *catManagementHandler) List(c *fiber.Ctx) error {
	queries := c.Queries()

	limit := queries["limit"]
	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		return invalidInput(c, err)
	}

	offset := queries["offset"]
	offsetNum, err := strconv.Atoi(offset)
	if err != nil {
		return invalidInput(c, err)
	}

	req := listCatsRequest{
		Id:         queries["id"],
		Limit: 		  limitNum,
		Offset: 	  offsetNum,
		Race:       queries["race"],
		Sex:        queries["sex"],
		HasMatched: queries["hasMatched"],
		AgeInMonth: queries["ageInMonth"],
		Owned:      queries["owned"],
		Name:       queries["search"],
	}

	if err := validate.Struct(req); err != nil {
		return invalidInput(c, err)
	}

	ownedBool, _ := strconv.ParseBool(req.Owned)

	getReq := domain.GetCatsRequest{
		Id:         req.Id,
		Limit:      req.Limit,
		Offset:     req.Offset,
		Race:       req.Race,
		Sex:        req.Sex,
		AgeInMonth: req.AgeInMonth,
		Name:       req.Name,
	}

	if ownedBool {
		getReq.UserId = "1"
	}

	cats, err := h.svc.List(&getReq)

	if err != nil {
		return serverError(c, fiber.StatusInternalServerError, "", err)
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    cats,
	})
}

func (h *catManagementHandler) Update(c *fiber.Ctx) error {
	req := new(createUpdateCatRequest)
	if err := c.BodyParser(req); err != nil {
		return failedToParseInput(c, err)
	}

	if err := validate.Struct(req); err != nil {
		return invalidInput(c, err)
	}

	cat := domain.Cat{
		Id:          c.Params("id"),
		Name:        req.Name,
		Race:        req.Race,
		Sex:         req.Sex,
		AgeInMonth:  req.AgeInMonth,
		Description: req.Description,
		ImageUrls:   req.ImageUrls,
		UserId:      "1",
	}

	updatedRecord, err := h.svc.Update(&cat)

	if err != nil {
		if err == domain.ErrNotFound {
			return serverError(c, fiber.StatusNotFound, "", err)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error",
			"errors": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"id": updatedRecord.Id,
			"updatedAt": updatedRecord.UpdatedAt,
		},
	})
}

func (h *catManagementHandler) Delete(c *fiber.Ctx) error {
	catId, deletedAt, err := h.svc.Delete("1", c.Params("id"))
	if err != nil {
		if err == domain.ErrNotFound {
			return serverError(c, fiber.StatusNotFound, "", err)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error",
			"errors": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"id": catId,
			"deletedAt": deletedAt.Format(iso8601),
		},
	})
}