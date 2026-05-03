package routes

import (
	"pm-gin/controllers"
	"pm-gin/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes sets up all authentication-related routes.
func UserRoutes(router *gin.Engine) {
	auth := router.Group("/api/User")
	{
		// Public routes
		auth.POST("/createuser", controllers.RegisterUser)

		// Protected routes
		auth.GET("/userprofile", middleware.AuthMiddleware(), controllers.GetUserProfile)
	}
}
