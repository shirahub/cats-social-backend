package api

import (
	"app/config"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=5,max=15"`
}

func getToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	return t, err
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	input := new(loginRequest)

	if err := c.BodyParser(input); err != nil {
		return failedToParseInput(c, err)
	}

	if err := validate.Struct(input); err != nil {
		return invalidInput(c, err)
	}

	user, err := h.repo.FindByEmail(c.Context(), input.Email)
	if err != nil {
		return serverError(c, err)
	}

	if !CheckPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": nil})
	}

	t, err := getToken(user.Id)
	if err != nil {
		return serverError(c, err)
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Success login",
		"data": fiber.Map{
			"email": user.Email,
			"name": user.Name,
			"accessToken": t,
		},
	})
}
