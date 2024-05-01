package api

import (
	"app/domain"
	"app/port"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

type catManagementHandler struct {
	svc port.CatManagementService
}

func NewCatManagementHandler(svc port.CatManagementService) *catManagementHandler {
	return &catManagementHandler{svc}
}
/*
Persian"
			- "Maine Coon"
			- "Siamese"
			- "Ragdoll"
			- "Bengal"
			- "Sphynx"
			- "British Shorthair"
			- "Abyssinian"
			- "Scottish Fold"
			- "Birman
*/
type createUpdateCatRequest struct {
	Name        string   `validate:"min=1,max=30"`
	Race        string   `validate:"oneof='Maine Coon' Siamese Ragdoll"`
	Sex         string   `validate:"oneof=male female"`
	AgeInMonth  int      `validate:"min=1,max=120082"`
	Description string   `validate:"min=1,max=200"`
	ImageUrls   []string `validate:"min=1,max=10"`
}

func (h *catManagementHandler) Create(c *fiber.Ctx) error {
	req := new(createUpdateCatRequest)
	if err := c.BodyParser(req); err != nil {
		return failedToParseInput(c, err)
	}

	validate := validator.New()
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
	h.svc.Create(&cat)

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
	req := new(createUpdateCatRequest)
	if err := c.BodyParser(req); err != nil {
		return failedToParseInput(c, err)
	}

	validate := validator.New()
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

	h.svc.Update(&cat)

	return c.JSON(fiber.Map{
		"status": "success",
	})
}

func (h *catManagementHandler) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
	})
}