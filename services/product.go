package services

import (
	"e-commerce/config"
	"e-commerce/dto"
	"e-commerce/models"
)

func CreateProduct(product dto.CreateProductDto, userId uint) (string, error) {
	newProduct := models.Product{
		Title:       product.Title,
		Price:       product.Price,
		Description: product.Description,
		CreatedBy:   userId,
	}

	if err := config.DB.Create(&newProduct).Error; err != nil {
		return "", err
	}

	return "Post created successfully", nil
}

func DeleteProduct(productId string) (string, error) {
	if err := config.DB.Where("id = ?", productId).Delete(&models.Product{}).Error; err != nil {
		return "", err
	}

	return "Post deleted successfully", nil
}
