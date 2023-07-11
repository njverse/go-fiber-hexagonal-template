package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/njverse/go-fiber-hexagonal-template/repository"
	"github.com/njverse/go-fiber-hexagonal-template/service"
)

// FiberHandler provide handlers in this app.
func FiberHandler(a *fiber.App, store *sqlx.DB) {
	todoRepositoryDB := repository.NewTodoRepositoryPostgres(store)
	todoService := service.NewTodoService(todoRepositoryDB)
	todoHandler := NewTodoHandler(todoService)

	// group request `/api` to apiGroup variable
	apiGroup := a.Group("/api")

	// group request `/api/v1` to v1Group variable
	v1Group := apiGroup.Group("/v1")
	v1Group.Get("/todos", todoHandler.GetTodos)
	v1Group.Post("/todos", todoHandler.NewTodo)
}
