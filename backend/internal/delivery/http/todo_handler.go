package http

import (
	"strconv"

	"github.com/hamwiwatsapon/todo-projects/backend/internal/domain"

	"github.com/gofiber/fiber/v3"
)

type TodoHandler struct {
	todoUsecase domain.TodoUsecase
}

func NewTodoHandler(app *fiber.App, usecase domain.TodoUsecase) {
	handler := &TodoHandler{
		todoUsecase: usecase,
	}

	app.Post("/todos", handler.CreateTodo)
	app.Get("/todos", handler.GetAllTodos)
	app.Get("/todos/:id", handler.GetTodo)
	app.Put("/todos/:id", handler.UpdateTodo)
	app.Delete("/todos/:id", handler.DeleteTodo)
}

func (h *TodoHandler) CreateTodo(c fiber.Ctx) error {
	todo := new(domain.Todo)
	if err := c.Bind().Body(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.todoUsecase.Create(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func (h *TodoHandler) GetAllTodos(c fiber.Ctx) error {
	todos, err := h.todoUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func (h *TodoHandler) GetTodo(c fiber.Ctx) error {
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

	return c.JSON(todo)
}

func (h *TodoHandler) UpdateTodo(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	todo := new(domain.Todo)
	if err := c.Bind().Body(todo); err != nil {
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

func (h *TodoHandler) DeleteTodo(c fiber.Ctx) error {
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
