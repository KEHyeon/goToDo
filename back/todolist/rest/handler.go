package rest

import (
	"go-api/database"
	"go-api/todolist/dblayer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db dblayer.DBLayer
}

func (h *Handler) GetAllTodo(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "dsn 오류"})
		return
	}
	todoList, err := h.db.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todoList)

	
}
type HandlerInterface interface {
	GetAllTodo(c *gin.Context)
	GetContent(c *gin.Context)
}

func NewHandler() (HandlerInterface, error) {
	dsn := database.DataSource

	db, err := dblayer.NewORM("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}