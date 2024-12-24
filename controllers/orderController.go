package controllers

import (
	"go-ecommerce-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetOrderStatus fetches the status of an order
func GetOrderStatus(c *gin.Context) {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	order, err := services.GetOrderStatus(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

