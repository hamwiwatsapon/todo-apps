package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	_ "github.com/hamwiwatsapon/todo-projects/backend/docs"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/delivery/http"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/repository"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/usecase"
	"github.com/hamwiwatsapon/todo-projects/backend/package/database"
)

// @title User API by Fiber and Swagger
// @version 1.0
// @description API user management Server by Fiber | Doc by Swagger.

// @contact.name wiwatsapon
// @contact.email hamlert33@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:3000
// @BasePath /
func main() {
	// Initialize database
	db, err := database.NewSQLiteDB()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository
	todoRepo := repository.NewTodoRepository(db)

	// Initialize usecase
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	// Initialize Fiber app
	app := fiber.New()

	// Add Logger middleware
	app.Use(logger.New())

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Configure Swagger
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	app.Get("/healthcheck", healthCheck)

	// Initialize handlers
	http.NewTodoHandler(app, todoUsecase)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// HealthCheckResponse represents the response structure
type HealthCheckResponse struct {
	Message string `json:"message" example:"OK"`
}

// HealthCheck
// @Summary Show the status of server.
// @Description Get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} HealthCheckResponse
// @Router /healthcheck [get]
func healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OK",
	})
}
