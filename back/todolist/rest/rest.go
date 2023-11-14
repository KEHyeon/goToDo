package rest

import "github.com/gin-gonic/gin"

func AddContentRoutes(rg *gin.RouterGroup) {
	content := rg.Group("/todo")
	h, _ := NewHandler()
	content.GET("/list", h.GetContent)
	content.GET("/:id", h.GetContent)
}