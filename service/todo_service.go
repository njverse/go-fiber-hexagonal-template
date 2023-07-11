package service

import (
	"github.com/njverse/go-fiber-hexagonal-template/repository"
)

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return todoService{todoRepo: todoRepo}
}

func (s todoService) NewTodo(request NewTodoRequest) (*TodoResponse, error) {
	todo := repository.Todo{
		Type:   request.Type,
		Name:   request.Name,
		Status: 1,
	}

	newTodo, err := s.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	res := TodoResponse{
		ID:     newTodo.ID,
		Type:   newTodo.Type,
		Name:   newTodo.Name,
		Status: newTodo.Status,
	}

	return &res, err
}

func (s todoService) GetTodos(request GetTodosRequest) ([]TodoResponse, error) {
	todo := repository.Todo{
		Type:   request.Type,
		Name:   request.Name,
		Status: request.Status,
	}

	todos, err := s.todoRepo.GetAll(todo)
	if err != nil {
		return nil, err
	}

	res := []TodoResponse{}
	for _, todo := range todos {
		res = append(res, TodoResponse{
			ID:     todo.ID,
			Type:   todo.Type,
			Name:   todo.Name,
			Status: todo.Status,
		})
	}

	return res, err
}
