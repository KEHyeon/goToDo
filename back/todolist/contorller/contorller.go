package controller

import "github.com/gin-gonic/gin"

func GetTodoRoutes(rg *gin.RouterGroup) {
	content := rg.Group("/todo")
	h, _ := GetHandler()

	content.GET("", h.GetAllTodo)
}