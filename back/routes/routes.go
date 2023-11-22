package routes

import (
	authController "go-api/auth/controller"
	todoController "go-api/todolist/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(address string) error {
	router := gin.Default()
	router.Use(cors.Default())
	v1 := router.Group("/")
	todoController.GetTodoRoutes(v1)
	authController.GetAuthRoutes(v1)
	return router.Run(address)
}