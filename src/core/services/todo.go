package services

import "github.com/qtpeters/pendiente/data"

func NewTodoService() *TodoService {
	return &TodoService{}
}

type TodoService struct{}

func (t *TodoService) CreateTodo(todo data.Todo) error {
	return nil
}

func (t *TodoService) UpdateTodo(todo data.Todo) error {
	return nil
}

func (t *TodoService) DeleteTodo(id string) error {
	return nil
}
