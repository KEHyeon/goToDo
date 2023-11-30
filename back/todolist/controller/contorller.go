package controller

import "github.com/gin-gonic/gin"

func GetTodoRoutes(rg *gin.RouterGroup) {
	todo := rg.Group("/todo")
	h, _ := GetHandler()
	
	todo.GET("/:date", h.GetAllTodoWithDate)
	todo.GET("/detail/:id", h.GetOneTodo)
	todo.POST("", h.CreateTodo)
	todo.POST("/toggle/:id", h.ToggleTodo)
}