package routes

import (
	"gin_project/controllers"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up the routes for the application
func InitializeRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)      // Create a new user
		userGroup.GET("/", controllers.GetAllUsers)      // Get all users
		userGroup.GET("/:id", controllers.GetUser)       // Get a user by ID
		userGroup.PUT("/:id", controllers.UpdateUser)    // Update a user by ID
		userGroup.DELETE("/:id", controllers.DeleteUser) // Delete a user by ID
	}
}
