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
type createCatRequest struct {
	Name        string   `validate:"min=1,max=30"`
	Race        string   `validate:"oneof='maine coon' siamese ragdoll"`
	Sex         string   `validate:"oneof=male female"`
	AgeInMonth  int      `validate:"min=1,max=120082"`
	Description string   `validate:"min=1,max=200"`
	ImageUrls   []string `validate:"min=1,max=10"`
}

func (h *catManagementHandler) Create(c *fiber.Ctx) error {
	cat := new(createCatRequest)
	if err := c.BodyParser(cat); err != nil {
		return failedToParseInput(c, err)
	}

	validate := validator.New()
	if err := validate.Struct(cat); err != nil {
		return invalidInput(c, err)
	}

	req := domain.CreateCatRequest{
		Name:        cat.Name,
		Race: 			 cat.Race,
		Sex:         cat.Sex,
		AgeInMonth:  cat.AgeInMonth,
		Description: cat.Description,
		ImageUrls:   cat.ImageUrls,
	}
	h.svc.Create(&req)

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