package routes

import (
	controller "go-api/todolist/contorller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(address string) error {
	router := gin.Default()
	router.Use(cors.Default())
	v1 := router.Group("/")
	controller.GetTodoRoutes(v1)
	return router.Run(address)
}