package routes

import (
	"e-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", controllers.Register)
	rg.POST("/login", controllers.Login)
}
