package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/delivery/http"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/repository"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/usecase"
	"github.com/hamwiwatsapon/todo-projects/backend/package/database"
)

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
	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "It's work!!!",
		})
	})

	// Initialize handlers
	http.NewTodoHandler(app, todoUsecase)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
