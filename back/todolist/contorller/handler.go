package controller

import (
	"go-api/database"
	"go-api/todolist/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	todoService service.TodoService
}

type HandlerInterface interface {
	GetAllTodo(c *gin.Context)
}


func GetHandler() (HandlerInterface, error) {
	dsn := database.DataSource()

	todoService, err := service.GetTodoService("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		todoService: todoService,
	}, nil
}


func (h *Handler) GetAllTodo(c *gin.Context) {
	todoList, err := h.todoService.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todoList)
}