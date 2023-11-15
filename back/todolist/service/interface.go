package service

import "go-api/todolist/models"

type TodoService interface {
	GetAllTodo() ([]models.Todo, error)
}