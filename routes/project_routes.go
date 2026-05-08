package routes

import (
	"pm-gin/controllers"
	"pm-gin/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes sets up all authentication-related routes.
func ProjectRoutes(router *gin.Engine) {
	auth := router.Group("/api/Project")
	{
		// Public routes
		//auth.POST("/createuser", controllers.RegisterUser)

		// Protected routes
		auth.GET("/getprojectbyid/:projectId", middleware.AuthMiddleware(), controllers.GetProjectById)
	}
}
