package service

import "go-api/todolist/models"

type TodoService interface {
	GetAllTodo() ([]models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
	ToggleTodo(id uint) (string, error)
	GetOneTodo(id uint) (models.Todo, error)
}