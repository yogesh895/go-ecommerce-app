package routes

import (
	"go-ecommerce-app/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Product routes
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductByID)

	// Order routes
	r.GET("/orders/:id", controllers.GetOrderStatus)
}

