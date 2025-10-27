package main

import (
	"e-commerce/config"
	"e-commerce/middlewares"
	"e-commerce/models"
	"e-commerce/routes"
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

	// Routes
	auth := router.Group("/auth")
	routes.AuthRoutes(auth)

	cart := router.Group("/cart", middlewares.CheckAuth)
	routes.CartRoutes(cart)

	product := router.Group("/product", middlewares.CheckAuth)
	routes.ProductRoutes(product)

	// Run Server
	port := os.Getenv("PORT")

	if os.Getenv("PORT") != "" {
		port = ":8080"
	}

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on port: " + port)
}
