package routers

import (
	"github.com/gin-gonic/gin"
	"practice_gin.com/internal/api/handlers"
)

func ProductsRouter(server *gin.Engine) {
	server.GET("/", handlers.HomePage)

	
	server.GET("/products/:id", handlers.GetProductByID)
	server.GET("/products", handlers.GetProducts)
	server.GET("/users/account", handlers.UserAccount)
	server.GET("/catalog", handlers.CatalogHandler)
	server.GET("/cart", handlers.CartHandler)
	server.GET("/admins", handlers.AdminsHandler)
	
	server.POST("/products", handlers.HandlerProductsForm)
	server.POST("/products/purchase", handlers.PurchaseProducts)
	server.POST("/", handlers.GetProductHomePage)
	server.POST("/products/add", handlers.AddProducts)
	server.POST("/products/add/categories", handlers.AddCategories)
	server.POST("/products/add/brands", handlers.AddBrands)
}
