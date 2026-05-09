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
		// Protected routes
		auth.GET("/getallprojects", middleware.AuthMiddleware(), controllers.GetAllProjects)
		auth.GET("/getprojectbyid/:projectId", middleware.AuthMiddleware(), controllers.GetProjectById)
		auth.POST("/createproject", middleware.AuthMiddleware(), controllers.CreateProject)
	}
}
