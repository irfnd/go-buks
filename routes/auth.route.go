package routes

import (
	"go-buks/handlers"
	"go-buks/middlewares"
	"go-buks/models"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(router fiber.Router) {
	r := router.Group("/auth")
	r.Post("/register", middlewares.ValidateBody[models.Register], handlers.Register)
	r.Post("/login", middlewares.ValidateBody[models.Login], handlers.Login)
}