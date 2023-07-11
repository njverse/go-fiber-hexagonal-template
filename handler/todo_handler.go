package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/njverse/go-fiber-hexagonal-template/service"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoSrv service.TodoService) TodoHandler {
	return todoHandler{todoService: todoSrv}
}

func (h todoHandler) GetTodos(c *fiber.Ctx) error {
	var req service.GetTodosRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	todos, err := h.todoService.GetTodos(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    todos,
	})
}

func (h todoHandler) NewTodo(c *fiber.Ctx) error {
	var req service.NewTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	todos, err := h.todoService.NewTodo(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    todos,
	})
}
