package routes

import (
	"go-buks/handlers"
	"go-buks/middlewares"
	"go-buks/models"

	"github.com/gofiber/fiber/v2"
)

func BooksRoute(route fiber.Router) {
	r := route.Group("/books")
	r.Get("/", middlewares.Auth, middlewares.ValidateQuery[models.QueryBook], handlers.BrowseBooks)
	r.Post("/", middlewares.Auth, middlewares.ValidateBody[models.CreateBook], handlers.CreateBook)
	r.Get("/:id", middlewares.Auth, middlewares.ValidateParams[models.ParamsBook], handlers.GetBook)
	r.Patch("/:id", middlewares.Auth, middlewares.ValidateParams[models.ParamsBook], middlewares.ValidateBody[models.UpdateBook], handlers.UpdateBook)
	r.Delete("/:id", middlewares.Auth, middlewares.ValidateParams[models.ParamsBook], handlers.DeleteBook)
}