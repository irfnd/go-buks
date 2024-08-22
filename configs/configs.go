package configs

import (
	"fmt"
	"go-buks/models"
	"go-buks/utils"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var validate = utils.ValidatorNew()
var Env *models.Config

func LoadEnv() {
	_ = godotenv.Load()

  Env = &models.Config {
  	APP_ENV:			getVariable("APP_ENV", "production"),
    PORT: 				getVariable("PORT", "8080"),
		HOST: 				getVariable("HOST", "localhost"),
    DATABASE_URI: getVariable("DATABASE_URI", ""),
		JWT_KEY:			getVariable("JWT_KEY", ""),
		JWT_EXP:			getVariable("JWT_EXP", ""),
  }
	
	validateEnv(Env)
}

func getVariable(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func validateEnv(env *models.Config) {
	if err := validate.Struct(env); err != nil {
		rsf := reflect.TypeOf(env).Elem()
		var errorMessages []string 
		for _, err := range err.(validator.ValidationErrors) {
			field, found := rsf.FieldByName(err.Field())
			if found {
				getMessage := fmt.Sprintf("- %s", field.Tag.Get("errormsg"))
				errorMessages = append(errorMessages, getMessage)
				} else {
				getMessage := fmt.Sprintf("%s Doesn't satisfy the `%s` constraint", err.Field(), err.Tag())
				errorMessages = append(errorMessages, getMessage)
			}
		}
		panicMessage := strings.Join(errorMessages, "\n")
		panic(fmt.Sprintf("Config validation errors:\n%s", panicMessage))
	}
}