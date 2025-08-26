package routers

import (
	"github.com/gin-gonic/gin"
	"practice_gin.com/internal/api/handlers"
)

func ProductsRouter(server *gin.Engine) {
	server.GET("/products/:id", handlers.GetProducts)
	server.POST("/products", handlers.PostProducts)
}
