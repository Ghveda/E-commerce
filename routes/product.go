package routes

import (
	"e-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {
	rg.GET("/list", controllers.GetProducts)
	rg.GET("/:id", controllers.GetProduct)
	rg.POST("/create", controllers.CreateProduct)
	rg.DELETE("/:productId", controllers.DeleteProduct)
}
