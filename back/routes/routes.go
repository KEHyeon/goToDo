package routes

import (
	todo_rest "go-api/todolist/rest"

	"github.com/gin-gonic/gin"
)

func Run(address string) error {
	router := gin.Default()
	v1 := router.Group("/")
	todo_rest.AddContentRoutes(v1)
	return router.Run(address)
}