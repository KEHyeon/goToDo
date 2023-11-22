package service

import "go-api/auth/models"

type AuthService interface {
	Signup(user models.User) (models.User, error)
	Signin(email string, password string) (string, error)
}