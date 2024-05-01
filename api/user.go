package api

import (
	"app/domain"
	"app/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
)

type userHandler struct {
	repo *repository.UserRepo
}

func NewUserHandler(repo *repository.UserRepo) *userHandler {
	return &userHandler{repo}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

type registerRequest struct {
	Email    string `validate:"required,email"`
	Name     string `validate:"min=5,max=50"`
	Password string `validate:"min=5,max=15"`
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	user := new(registerRequest)
	if err := c.BodyParser(user); err != nil {
		return failedToParseInput(c, err)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return invalidInput(c, err)
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "errors": err.Error()})
	}

	err = h.repo.Create(domain.User{Email: user.Email, Name: user.Name, Password: hash})
	if err != nil {
		log.Error(err)
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}
