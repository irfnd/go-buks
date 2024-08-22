package utils

import (
	"go-buks/models"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(code int, message string, reason interface{}) *models.CustomError {
	return &models.CustomError{
		Statuscode: code,
		Message:		message,
		Reason:			reason,
	}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	response := &models.CustomError{
		Statuscode: 	fiber.StatusInternalServerError,
		Message:    	"Internal Server Error",
		Reason:				[]string{},
	}

	if customErr, ok := err.(*models.CustomError); ok {
		response.Statuscode	= customErr.Statuscode
		response.Message		= customErr.Message
		response.Reason			= customErr.Reason
	}
	
	if fiberErr, ok := err.(*fiber.Error); ok {
		response.Statuscode = fiberErr.Code
		response.Message		= fiberErr.Message
		response.Reason			= []string{}
	}

	return c.Status(response.Statuscode).JSON(response)
}