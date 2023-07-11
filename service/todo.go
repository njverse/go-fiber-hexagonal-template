package service

import "github.com/google/uuid"

type NewTodoRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetTodosRequest struct {
	Name   string `query:"name"`
	Type   string `query:"type"`
	Status int    `query:"status"`
}

type TodoResponse struct {
	ID     uuid.UUID `json:"id"`
	Type   string    `json:"type"`
	Name   string    `json:"name"`
	Status int       `json:"status"`
}

type TodoService interface {
	NewTodo(NewTodoRequest) (*TodoResponse, error)
	GetTodos(GetTodosRequest) ([]TodoResponse, error)
}
