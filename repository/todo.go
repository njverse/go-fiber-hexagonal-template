package repository

import "github.com/google/uuid"

type Todo struct {
	ID     uuid.UUID `db:"id"`
	Type   string    `db:"type"`
	Name   string    `db:"name"`
	Status int       `db:"status"`
}

type TodoRepository interface {
	Create(Todo) (*Todo, error)
	GetAll(Todo) ([]Todo, error)
}
