package services

import (
	"go-ecommerce-app/models"
	"errors"
)

// Mock product data
var products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Smartphone", Price: 699.99},
	{ID: 3, Name: "Headphones", Price: 199.99},
}

// GetAllProducts retrieves all products
func GetAllProducts() ([]models.Product, error) {
	return products, nil
}

// GetProductByID retrieves a single product by its ID
func GetProductByID(id int) (models.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

