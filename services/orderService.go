package services

import (
	"go-ecommerce-app/models"
	"errors"
)

// Mock order data
var orders = []models.Order{
	{ID: 1, ProductID: 1, Quantity: 1, Status: "Shipped"},
	{ID: 2, ProductID: 2, Quantity: 2, Status: "Delivered"},
}

// GetOrderStatus retrieves the status of an order by its ID
func GetOrderStatus(id int) (models.Order, error) {
	for _, order := range orders {
		if order.ID == id {
			return order, nil
		}
	}
	return models.Order{}, errors.New("order not found")
}

