package routes

import (
	"pm-gin/controllers"
	"pm-gin/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes sets up all authentication-related routes.
func MemberRoleRoutes(router *gin.Engine) {
	memberRole := router.Group("/api/MemberRole")
	{
		// member role routes
		memberRole.PUT("/updatememberrolebyprojectmemberid/:projectMemberId", middleware.AuthMiddleware(), controllers.UpdateMemberRole)
	}
}
