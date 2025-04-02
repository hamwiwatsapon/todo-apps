package http

import (
	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(app *fiber.App, usecase domain.AuthUsecase) {
	handler := &AuthHandler{
		authUsecase: usecase,
	}

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Post("/refreshToken", handler.GetRefreshToken)

	// Apply JWT middleware for routes that require authentication
	authenticated := app.Group("/auth", middleware.JwtMiddleware())
	authenticated.Delete("/delete/:email", handler.DeleteUser)
}

// Register.
// @Summary Register user.
// @Description Use for create user.
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.CreateUserDTO true "User object"
// @Success 201 {object} domain.User
// @Failure 400 {object} domain.ErrorResponse400 "Bad Request"
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Router /register [post]
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

// Login.
// @Summary Login user.
// @Description Use for Login.
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.LoginDTO true "User object"
// @Success 200 {object} domain.LoginReturn
// @Failure 400 {object} domain.ErrorResponse400 "Bad Request"
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Router /login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	user := new(domain.LoginDTO)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, refreshToken, err := h.authUsecase.Login(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Login successful",
		"token":         token,
		"refresh_token": refreshToken,
	})
}

// Refresh Token.
// @Summary Refresh Token.
// @Description Use for Get New Token.
// @Tags users
// @Accept json
// @Produce json
// @Param refreshToken body domain.RefreshTokenDTO true "token"
// @Success 200 {object} domain.RefreshTokenReturn
// @Failure 401 {object} domain.ErrorResponse401 "Unauthorized"
// @Failure 403 {object} domain.ErrorResponse403 "Forbidden"
// @Failure 400 {object} domain.ErrorResponse400 "Bad Request"
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Router /refreshToken [post]
func (h *AuthHandler) GetRefreshToken(c *fiber.Ctx) error {
	refreshTokenDTO := new(domain.RefreshTokenDTO)
	if err := c.BodyParser(refreshTokenDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	newToken, newRefreshToken, err := h.authUsecase.RefreshToken(refreshTokenDTO.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Token refreshed successfully",
		"token":         newToken,
		"refresh_token": newRefreshToken,
	})
}

// DeleteUser.
// @Summary Delete user.
// @Description Use for delete user by email.
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} domain.User
// @Failure 400 {object} domain.ErrorResponse400 "Bad Request"
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Router /auth/delete/{email} [delete]
func (h *AuthHandler) DeleteUser(c *fiber.Ctx) error {
	panic("unimplemented")
}
