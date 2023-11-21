package controller

import "github.com/gin-gonic/gin"

func GetTodoRoutes(rg *gin.RouterGroup) {
	todo := rg.Group("/todo")
	h, _ := GetHandler()

	todo.GET("", h.GetAllTodo)
	todo.POST("", h.CreateTodo)
}