package handlers

import (
	"errors"
	"go-buks/models"
	"go-buks/services"
	"go-buks/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register
//	@Summary	Register New User
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		data			body		models.Register	true	"User Data"
//	@Success	200				{object}	models.RegisterResponse
//	@Failure	400				{object}	models.RegisterResponseBadRequest
//	@Failure	500				{object}	models.CustomErrorResponse
//	@Router		/auth/register 	[post]
func Register(c *fiber.Ctx) error {
	body := c.Locals("body").(models.Register)
	
	results := &models.User{
		Fullname: body.Fullname,
		Email: body.Email,
		Password: services.GenerateHash(body.Password),
	}

	response := &models.RegisterResponse{
		Statuscode: fiber.StatusCreated,
		Message: "Registered Successfully",
		Data: results,
	}

	checkUser := services.GetUser(&models.User{}, "email = ?", body.Email).Error
	if !errors.Is(checkUser, gorm.ErrRecordNotFound) {
		return utils.ErrorResponse(fiber.StatusConflict, "Register Failed", "Email Already Exists")
	}
	if err := services.CreateUser(results).Error; err != nil {
		return utils.ErrorResponse(fiber.StatusConflict, "Register Failed", err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// Login
//	@Summary	Login New User
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		data			body		models.Login	true	"Login Credentials"
//	@Success	200				{object}	models.LoginResponse
//	@Failure	400				{object}	models.LoginResponseBadRequest
//	@Failure	500				{object}	models.CustomErrorResponse
//	@Router		/auth/login 	[post]
func Login(c *fiber.Ctx) error {
	body := c.Locals("body").(models.Login)

	results := &models.User{}
	
	checkUser := services.GetUser(results, "email = ?", body.Email).Error
	if errors.Is(checkUser, gorm.ErrRecordNotFound) {
		return utils.ErrorResponse(fiber.StatusConflict, "Login Failed",  "Invalid Email or Password")
	}
	if err := services.VerifyHash(results.Password, body.Password); err != nil {
		return utils.ErrorResponse(fiber.StatusConflict, "Login Failed",  "Invalid Email or Password")
	}
	
	token := services.GenerateJwt(&models.TokenPayload{
		ID: results.ID,
		Email: results.Email,
	})
	
	response := &models.LoginResponse{}
	response.Statuscode = fiber.StatusOK
	response.Message = "Logged in Successfully"
	response.Data.Token = token

	return c.Status(fiber.StatusOK).JSON(response)
}