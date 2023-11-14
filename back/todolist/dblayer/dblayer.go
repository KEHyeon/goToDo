package dblayer

import "go-api/todolist/models"

type DBLayer interface {
	GetAllTodo() ([]models.Todo, error)
}