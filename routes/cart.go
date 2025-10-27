package routes

import (
	"e-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoutes(rg *gin.RouterGroup) {
	rg.POST("/add", controllers.AddProductToCart)
	rg.DELETE("/delete", controllers.DeleteProductFromCart)
}
