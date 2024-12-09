package routes

import (
	"product-management/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProductByID)
	r.GET("/products", handlers.GetProducts)

	return r
}
