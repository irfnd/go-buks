package middlewares

import (
	"go-buks/services"
	"go-buks/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorization")
	if h == "" {
		return utils.ErrorResponse(fiber.StatusUnauthorized, "Unauthorized", "Token Required")
	}

	chunks := strings.Split(h, " ")
	if len(chunks) < 2 {
		return utils.ErrorResponse(fiber.StatusUnauthorized, "Unauthorized", "Token Required")
	}

	user, err := services.VerifyJwt(chunks[1])
	if err != nil {
		return utils.ErrorResponse(fiber.StatusUnauthorized, "Unauthorized", cases.Title(language.AmericanEnglish).String(err.Error()))
	}

	c.Locals("user", user)
	return c.Next()
}