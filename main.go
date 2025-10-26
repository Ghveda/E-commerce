package main

import (
	"e-commerce/config"
	"e-commerce/controllers"
	"e-commerce/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DbConnect()
	models.Migrate()
	
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	cart := router.Group("/cart")
	{
		cart.POST("/add", controllers.AddProductToCart)
		cart.DELETE("/delete", controllers.DeleteProductFromCart)
	}

	product := router.Group("/product")
	{
		product.GET("/list", controllers.GetProducts)
		product.GET("/:id", controllers.GetProduct)
		product.POST("/add", controllers.CreateProduct)
		product.DELETE("/delete", controllers.DeleteProduct)
	}

	port := os.Getenv("PORT")

	if os.Getenv("PORT") != "" {
		port = ":8080"
	}

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on port: " + port)
}
