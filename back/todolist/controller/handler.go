package controller

import (
	"fmt"
	"go-api/database"
	"go-api/todolist/models"
	"go-api/todolist/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	todoService service.TodoService
}

type HandlerInterface interface {
	GetAllTodoWithDate(c *gin.Context)
	GetOneTodo(c *gin.Context)
	ToggleTodo(c *gin.Context)
	CreateTodo(c *gin.Context)
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
func (h *Handler) GetOneTodo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	res, err  := h.todoService.GetOneTodo(uint(id))
	if(err != nil) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} 
	c.JSON(http.StatusOK, res)
}
func (h *Handler) ToggleTodo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	fmt.Print(id)
	res, err := h.todoService.ToggleTodo(uint(id))
	if(err != nil) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllTodoWithDate(c *gin.Context) {
	todoList, err := h.todoService.GetAllTodoWithDate(c.Param("date"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todoList)
}

func (h *Handler) CreateTodo(c *gin.Context) {
	var todoData models.Todo
	err := c.ShouldBindJSON(&todoData)
	loc, _ := time.LoadLocation("Asia/Seoul")
	todoData.CreatedAt = time.Now().In(loc).String()
	fmt.Println(todoData.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	todo,err := h.todoService.CreateTodo(todoData)
	if(err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}
