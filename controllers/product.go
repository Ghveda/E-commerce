package controllers

import (
	"e-commerce/dto"
	"e-commerce/services"
	"e-commerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func GetProduct(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func CreateProduct(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")

	isAdmin, user := utils.IsAdmin(currentUser)

	if !isAdmin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You dont have access to create a product"})
		return
	}

	var product dto.CreateProductDto

	if err := c.ShouldBindBodyWithJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := services.CreateProduct(product, user.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": message})
}

func DeleteProduct(c *gin.Context) {
	productId := c.Param("productId")
	currentUser, _ := c.Get("currentUser")

	if isAdmin, _ := utils.IsAdmin(currentUser); !isAdmin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You dont have access to create a product"})
		return
	}

	message, err := services.DeleteProduct(productId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
