package controllers

import (
	"e-commerce/dto"
	"e-commerce/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var authBody dto.AuthBody

	if err := c.ShouldBindBodyWithJSON(&authBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := services.Register(authBody.Username, authBody.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": message})
}

func Login(c *gin.Context) {
	var authBody dto.AuthBody

	if err := c.ShouldBindBodyWithJSON(&authBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Login(authBody.Username, authBody.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"token": token})
}
