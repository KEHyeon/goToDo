package service

import "go-api/auth/models"

type TodoService interface {
	Register(user models.User) (models.User, error)
	Login(email string, password string) (models.User, error)
}