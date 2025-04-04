package http

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"
	"github.com/hamwiwatsapon/todo-projects/backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoUsecase domain.TodoUsecase
}

func NewTodoHandler(app *fiber.App, usecase domain.TodoUsecase) {
	handler := &TodoHandler{
		todoUsecase: usecase,
	}

	authenticated := app.Group("/auth", middleware.JwtMiddleware())
	authenticated.Post("/todos", handler.CreateTodo)
	authenticated.Get("/todos", handler.GetAllTodos)
	authenticated.Get("/todos/:id", handler.GetTodo)
	authenticated.Put("/todos/:id", handler.UpdateTodo)
	authenticated.Delete("/todos/:id", handler.DeleteTodo)
}

// Create to do.
// @Summary Create to do.
// @Description Use for create to do.
// @Authrorization Bearer
// @Tags todo
// @Accept json
// @Produce json
// @Param todo body domain.CreateTodoDTO true "Todo object"
// @Success 201 {object} domain.Todo
// @Failure 400 {object} domain.ErrorResponse400 "Bad Request"
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Router /auth/todos [post]
func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	// Extract user claims from context
	userClaims := c.Locals("user").(jwt.MapClaims)

	// Extract userID from claims
	userID, ok := userClaims["user_id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user ID in token",
		})
	}

	todo := new(domain.CreateTodoDTO)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Set the userID in the DTO
	todo.UserID = userID

	createdTodo, err := h.todoUsecase.Create(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdTodo)
}

// Get all todos
// @Summary Get all todos
// @Description Retrieve all todos from the system
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {array} domain.Todo
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Security BearerAuth
// @Router /auth/todos [get]
func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	// Extract user claims from context
	userClaims := c.Locals("user").(jwt.MapClaims)

	// Extract userID from claims
	userID, ok := userClaims["user_id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user ID in token",
		})
	}

	todos, err := h.todoUsecase.GetAllByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

// Get todo by id
// @Summary Get todo by id
// @Description Retrieve todo by id from the system
// @Tags todo
// @Accept json
// @Param id path int true "Todo ID"
// @Produce json
// @Success 200 {object} domain.Todo
// @Failure 400 {object} domain.ErrorResponse400 "Bad Request"
// @Failure 404 {object} domain.ErrorResponse404 "Todo not found"
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Security BearerAuth
// @Router /auth/todos/{id} [get]
func (h *TodoHandler) GetTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	todo, err := h.todoUsecase.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

// Update todo
// @Summary Update todo
// @Description Update todo to the system
// @Tags todo
// @Accept json
// @Produce json
// @Param id path uint true "Todo ID"
// @Success 200 {object} domain.Todo
// @Failure 404 {object} domain.ErrorResponse404 "Todo not found"
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Security BearerAuth
// @Router /auth/todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	todo := new(domain.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	todo.ID = uint(id)
	if err := h.todoUsecase.Update(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(todo)
}

// Delete todo
// @Summary Delete todo
// @Description Delete todo to the system
// @Tags todo
// @Accept json
// @Produce json
// @Param id path uint true "Todo ID"
// @Success 200 {object} domain.Todo
// @Failure 400 {object} domain.ErrorResponse404 "Todo not found"
// @Failure 500 {object} domain.ErrorResponse500 "Internal Server Error"
// @Security BearerAuth
// @Router /auth/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := h.todoUsecase.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
