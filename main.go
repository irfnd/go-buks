package main

import (
	"fmt"
	"go-buks/configs"
	"go-buks/database"
	"go-buks/docs"
	"go-buks/models"
	"go-buks/routes"
	"go-buks/utils"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@securityDefinitions.apikey	BearerAuth
//	@name						Authorization
//	@in							header
//	@description				Type "Bearer" followed by a space and then your token
func main() {
	// Load Env
	configs.LoadEnv()

	// Database
	database.Connect(configs.Env.DATABASE_URI)
	database.Migrate(&models.User{}, &models.Book{})

	// Server
	port := fmt.Sprintf(":%s", configs.Env.PORT)
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: utils.ErrorHandler,
	})

	// Swagger
	docs.SwaggerInfo.Title = "Go-Buks API"
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.BasePath = "/"
	if configs.Env.APP_ENV == "production" {
		docs.SwaggerInfo.Host = fmt.Sprintf("%v", configs.Env.HOST)
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		docs.SwaggerInfo.Host = fmt.Sprintf("%v%v", configs.Env.HOST, port)
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	// Middlewares
	app.Use(logger.New(), helmet.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"statusCode": fiber.StatusOK,
			"message":    "Welcome to Go-buks API",
			"techs":      []string{"Fiber", "PostgreSQL"},
		})
	})
	app.Get("/docs/*", swagger.HandlerDefault)
	routes.Routes(app)

	// Running Server
	app.Listen(port)
}
