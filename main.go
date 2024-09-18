package main

import (
	"gin_project/database"
	"gin_project/routes"

	"github.com/gin-gonic/gin"
)

// Main function
func main() {

	// Initialize Gin router
	router := gin.Default()

	// Initialize the routes
	routes.InitializeRoutes(router)

	database.ConnectDB()

	router.SetTrustedProxies([]string{"127.0.0.1"}) // Only trust localhost

	// Start the Gin server on port 8080
	router.Run(":8080")
}
