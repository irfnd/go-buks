package middlewares

import (
	"fmt"
	"go-buks/models"
	"go-buks/utils"
	"reflect"
	"strings"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var validate = utils.ValidatorNew()

func ValidateBody[T interface{}] (c *fiber.Ctx) error {
	var body T
	
	if err := parseAndSetDefault(&body, c.BodyParser(&body)); err != nil {
		return utils.ErrorResponse(fiber.StatusBadRequest, "Parse Body Failed", err.Error())
	}

	if err := validate.Struct(body); err != nil {
		rsf := reflect.TypeOf(body)
		var errors []models.ValidationDetail
		for _, err := range err.(validator.ValidationErrors) {
			field, found := rsf.FieldByName(err.Field())
			message := func () string {
				if found {
					return field.Tag.Get("errormsg")
				} else {
					return fmt.Sprintf("%v Doesn't satisfy the `%v` constraint", strings.ToLower(err.Field()), err.Tag())
				}
			}()

			errorMap := models.ValidationDetail {
				Field: strings.ToLower(err.Field()),
				Value: fmt.Sprintf("%v", err.Value()),
				Message: cases.Title(language.AmericanEnglish).String(message),
			}
			errors = append(errors, errorMap)
		}
		return utils.ErrorResponse(fiber.StatusBadRequest, "Validate Body Failed", errors)
	}
	
	c.Locals("body", body)
	return c.Next()
}

func ValidateQuery[T interface{}] (c *fiber.Ctx) error {
	var queries T
	
	if err := parseAndSetDefault(&queries, c.QueryParser(&queries)); err != nil {
		return utils.ErrorResponse(fiber.StatusBadRequest, "Parse Query Failed", err.Error())
	}

	if err := validate.Struct(queries); err != nil {
		rsf := reflect.TypeOf(queries)
		var errors []models.ValidationDetail
		for _, err := range err.(validator.ValidationErrors) {
			field, found := rsf.FieldByName(err.Field())
			message := func () string {
				if found {
					return field.Tag.Get("errormsg")
				} else {
					return fmt.Sprintf("%v Doesn't satisfy the `%v` constraint", strings.ToLower(err.Field()), err.Tag())
				}
			}()
			
			errorMap := models.ValidationDetail {
				Field: strings.ToLower(err.Field()),
				Value: fmt.Sprintf("%v", err.Value()),
				Message: cases.Title(language.AmericanEnglish).String(message),
			}
			errors = append(errors, errorMap)
		}
		return utils.ErrorResponse(fiber.StatusBadRequest, "Validate Queries Failed", errors)
	}

	c.Locals("queries", queries)
	return c.Next()
}

func ValidateParams[T interface{}] (c *fiber.Ctx) error {
	var params T
	
	if err := c.ParamsParser(&params); err != nil {
		return utils.ErrorResponse(fiber.StatusBadRequest, "Parse Params Failed", err.Error())
	}
	
	if err := validate.Struct(params); err != nil {
		rsf := reflect.TypeOf(params)
		var errors []models.ValidationDetail
		for _, err := range err.(validator.ValidationErrors) {
			field, found := rsf.FieldByName(err.Field())
			message := func () string {
				if found {
					return field.Tag.Get("errormsg")
				} else {
					return fmt.Sprintf("%v Doesn't satisfy the `%v` constraint", strings.ToLower(err.Field()), err.Tag())
				}
			}()

			errorMap := models.ValidationDetail {
				Field: strings.ToLower(err.Field()),
				Value: fmt.Sprintf("%v", err.Value()),
				Message: cases.Title(language.AmericanEnglish).String(message),
			}
			errors = append(errors, errorMap)
		}
		return utils.ErrorResponse(fiber.StatusBadRequest, "Validate Params Failed", errors)
	}
	
	c.Locals("params", params)
	return c.Next()
}

func parseAndSetDefault[T interface{}](target *T, parser error) error {
	if err := parser; err != nil {
		return err
	}
	if err := defaults.Set(target); err != nil {
		return err
	}
	return nil
}