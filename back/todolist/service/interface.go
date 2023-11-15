package service

import "go-api/todolist/models"

type TodoService interface {
	GetAllTodo() ([]models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
}