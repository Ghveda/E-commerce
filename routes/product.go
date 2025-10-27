package routes

import (
	"e-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {
	rg.POST("/add", controllers.AddProductToCart)
	rg.DELETE("/delete", controllers.DeleteProductFromCart)
}
