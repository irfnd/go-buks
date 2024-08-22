package handlers

import (
	"go-buks/models"
	"go-buks/services"
	"go-buks/utils"

	"github.com/gofiber/fiber/v2"
)

// BrowseBooks
//	@Summary	Browse Books by User ID
//	@Tags		Books
//	@Accept		json
//	@Produce	json
//	@Param		search	query		string	false	"Search book by title"
//	@Param		page	query		int		false	"Current page"			minimum(1)	default(1)
//	@Param		limit	query		int		false	"limit items per page"	minimum(1)	default(10)
//	@Success	200		{object}	models.QueryBookResponse
//	@Failure	400		{object}	models.QueryBookResponseBadRequest
//	@Failure	401		{object}	models.UnauthorizedResponse
//	@Failure	500		{object}	models.CustomErrorResponse
//	@Router		/books 	[get]
//	@Security	BearerAuth
func BrowseBooks(c *fiber.Ctx) error {
	queries := c.Locals("queries").(models.QueryBook)
	user := c.Locals("user").(*models.TokenPayload)

	results := &[]models.Book{}
	
	response := &models.QueryBookResponse{}
	response.Statuscode = fiber.StatusOK
	response.Message = "Retrieved Books Successfully"
	response.Data.Books = results
	response.Data.Pagination.Page = queries.Page
	response.Data.Pagination.Limit = queries.Limit

	if err := services.BrowseBooks(results, queries, user.ID).Error; err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, "Browse Books Failed", err.Error())
	}
	return c.JSON(response)
}

// GetBook
//	@Summary	Get a Book
//	@Tags		Books
//	@Accept		json
//	@Produce	json
//	@Param		id			path		string	true	"Book ID"
//	@Success	200			{object}	models.GetBookResponse
//	@Failure	400			{object}	models.GetBookResponseBadRequest
//	@Failure	401			{object}	models.UnauthorizedResponse
//	@Failure	500			{object}	models.CustomErrorResponse
//	@Router		/books/{id}	[get]
//	@Security	BearerAuth
func GetBook(c *fiber.Ctx) error {
	params := c.Locals("params").(models.ParamsBook)
	user := c.Locals("user").(*models.TokenPayload)

	results := &models.Book{}
	
	response := &models.GetBookResponse{
		Statuscode: fiber.StatusOK,
		Message: "Retrieved Book Successfully",
		Data: results,
	}
	
	if err := services.GetBook(results, params.ID, user.ID).Error; err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, "Get Book Failed", err.Error())
	}
	return c.JSON(response)
}

// CreateBook
//	@Summary	Create a book
//	@Tags		Books
//	@Accept		json
//	@Produce	json
//	@Param		data	body		models.CreateBook	true	"New Book Data"
//	@Success	201		{object}	models.CreateBookResponse
//	@Failure	400		{object}	models.CreateBookResponseBadRequest
//	@Failure	401		{object}	models.UnauthorizedResponse
//	@Failure	500		{object}	models.CustomErrorResponse
//	@Router		/books 	[post]
//	@Security	BearerAuth
func CreateBook(c *fiber.Ctx) error {
	body := c.Locals("body").(models.CreateBook)
	user := c.Locals("user").(*models.TokenPayload)

	results := &models.Book{
		Title: body.Title,
		Author: body.Author,
		Publisher: body.Publisher,
		Year: body.Year,
		UserID: user.ID,
	}
	
	response := &models.CreateBookResponse{
		Statuscode: fiber.StatusCreated,
		Message: "Created Book Successfully",
		Data: results,
	}

	if err := services.CreateBook(results).Error; err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, "Create Book Failed", err.Error())
	}
	return c.Status(response.Statuscode).JSON(response)
}

// UpdateBook
//	@Summary	Update a Book
//	@Tags		Books
//	@Accept		json
//	@Produce	json
//	@Param		id			path		string				true	"Book ID"
//	@Param		data		body		models.UpdateBook	true	"New Data"
//	@Success	200			{object}	models.UpdateBookResponse
//	@Failure	400			{object}	models.UpdateBookResponseBadRequest
//	@Failure	401			{object}	models.UnauthorizedResponse
//	@Failure	500			{object}	models.CustomErrorResponse
//	@Router		/books/{id}	[patch]
//	@Security	BearerAuth
func UpdateBook(c *fiber.Ctx) error {
	params := c.Locals("params").(models.ParamsBook)
	body := c.Locals("body").(models.UpdateBook)
	user := c.Locals("user").(*models.TokenPayload)

	results := &models.Book{
		Title: body.Title,
		Author: body.Author,
		Publisher: body.Publisher,
		Year: body.Year,
	}
	
	response := &models.UpdateBookResponse{
		Statuscode: fiber.StatusOK,
		Message: "Updated Book Successfully",
		Data: results,
	}

	if err := services.UpdateBook(results, params.ID, user.ID).Error; err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, "Update Book Failed", err.Error())
	}
	return c.JSON(response)
}

// DeleteBook
//	@Summary	Delete a Book
//	@Tags			Books
//	@Accept		json
//	@Produce	json
//	@Param		id					path			string			true	"Book ID"
//	@Success	200					{object}	models.DeleteBookResponse
//	@Failure	400					{object}	models.DeleteBookResponseBadRequest
//	@Failure	401			{object}	models.UnauthorizedResponse
//	@Failure	500			{object}	models.CustomErrorResponse
//	@Router		/books/{id}	[delete]
//	@Security	BearerAuth
func DeleteBook(c *fiber.Ctx) error {
	params := c.Locals("params").(models.ParamsBook)
	user := c.Locals("user").(*models.TokenPayload)

	results := &models.Book{}
	
	response := &models.DeleteBookResponse{
		Statuscode: fiber.StatusOK,
		Message: "Deleted Book Successfully",
		Data: results,
	}

	if err := services.DeleteBook(results, params.ID, user.ID).Error; err != nil {
		return utils.ErrorResponse(fiber.StatusInternalServerError, "Delete Book Failed", err.Error())
	}
	return c.JSON(response)
}