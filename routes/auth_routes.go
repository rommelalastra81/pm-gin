package routes

import (
	"pm-gin/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes sets up all authentication-related routes.
func RegisterAuthRoutes(router *gin.Engine) {
	//auth := router.Group("/api/auth")
	auth := router.Group("/api")
	{
		// Public routes
		auth.POST("/login", controllers.Login)
	}
}
