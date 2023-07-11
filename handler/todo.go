package handler

import "github.com/gofiber/fiber/v2"

type TodoHandler interface {
	GetTodos(*fiber.Ctx) error
	NewTodo(*fiber.Ctx) error
}
