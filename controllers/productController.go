package controllers

import (
	"go-ecommerce-app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllProducts fetches the list of products
func GetAllProducts(c *gin.Context) {
	products, err := services.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID fetches a single product by its ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := services.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
