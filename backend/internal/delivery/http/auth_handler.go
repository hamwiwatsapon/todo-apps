package http

import (
	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(app *fiber.App, usecase domain.AuthUsecase) {
	handler := &AuthHandler{
		authUsecase: usecase,
	}

	app.Post("/auth/register", handler.Register)
	// app.Post("/auth/login", handler.Login)
	// app.Delete("/auth/delete/:email", handler.DeleteUser)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	user := new(domain.CreateUserDTO)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdUser, err := h.authUsecase.CreateUser(user)
	// Check if the email is already taken
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    createdUser.Email,
	})
}
