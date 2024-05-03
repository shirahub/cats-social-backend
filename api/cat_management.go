package api

import (
	"app/domain"
	"app/port"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
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
	ImageUrls   []string `validate:"min=1,max=10,dive,url"`
}

type listCatsRequest struct {
	Id         string
	Limit      int    `validate:"gte=0"`
	Offset     int    `validate:"gte=0"`
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
		UserId:      getUserId(c),
	}
	newRecord, err := h.svc.Create(c.Context(), &cat)
	if err != nil {
		return serverError(c, err)
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

	limitNum, err := strconv.Atoi(queries["limit"])
	if err != nil {
		limitNum = 5
	}

	offsetNum, err := strconv.Atoi(queries["offset"])
	if err != nil {
		offsetNum = 0
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

	getReq := domain.GetCatsRequest{
		Id:         req.Id,
		Limit:      req.Limit,
		Offset:     req.Offset,
		Race:       req.Race,
		Sex:        req.Sex,
		AgeInMonth: req.AgeInMonth,
		Name:       req.Name,
	}

	ownedBool, _ := strconv.ParseBool(req.Owned)
	if ownedBool {
		getReq.UserId = getUserId(c)
	}

	matchedBool, err := strconv.ParseBool(req.HasMatched)
	if err == nil {
		getReq.HasMatched = &matchedBool
	}

	if err := validate.Struct(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range errs {
				switch fe.StructField() {
				case "Limit":
					getReq.Limit = 5
				case "Offset":
					getReq.Offset = 0
				case "AgeInMonth":
					getReq.AgeInMonth = ""
				case "Sex":
					getReq.Sex = ""
				case "HasMatched":
					getReq.HasMatched = nil
				}
			}
		}
	}

	cats, err := h.svc.List(c.Context(), &getReq)

	if err != nil {
		return serverError(c, err)
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
		UserId:      getUserId(c),
	}

	updatedRecord, err := h.svc.Update(c.Context(), &cat)
	if err != nil {
		return serverError(c, err)
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
	catId, deletedAt, err := h.svc.Delete(c.Context(), c.Params("id"), getUserId(c))
	if err != nil {
		if err == domain.ErrNotFound {
			return serverError(c, err)
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