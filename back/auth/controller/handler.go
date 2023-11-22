package controller

import (
	"go-api/auth/models"
	"go-api/auth/service"
	"go-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService service.AuthService
}

type HandlerInterface interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}

func GetHandler() (HandlerInterface, error) {
	dsn := database.DataSource()
	authService, err := service.GetAuthService("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		authService: authService,
	}, nil
}

func (h *Handler) Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user, err := h.authService.Signup(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Signin(c *gin.Context) {}