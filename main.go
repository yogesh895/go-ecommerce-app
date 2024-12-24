package main

import (
	"go-ecommerce-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}

