package routes

import (
	"pm-gin/controllers"
	"pm-gin/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes sets up all authentication-related routes.
func ProjectMemberRoutes(router *gin.Engine) {
	auth := router.Group("/api/ProjectMember")
	{
		// Protected routes
		auth.POST("/addprojectmembers/:projectId", middleware.AuthMiddleware(), controllers.AddProjectMembers)
		auth.GET("/getmembersbyprojectid/:projectId/:page/:pageSize", middleware.AuthMiddleware(), controllers.GetMembersByProjectId)
		auth.GET("/getusersnotonproject/:projectId", middleware.AuthMiddleware(), controllers.GetUsersNotOnProject)
	}
}
