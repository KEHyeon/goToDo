package controller

import "github.com/gin-gonic/gin"

func GetAuthRoutes(rg *gin.RouterGroup) {
	todo := rg.Group("/auth")
	h, _ := GetHandler()
	
	todo.POST("/signup", h.Signup)
	todo.POST("/signin", h.Signin)
}