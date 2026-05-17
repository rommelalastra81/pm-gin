package routes

import (
	"pm-gin/controllers"
	"pm-gin/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	auth := router.Group("/api/Task")
	{
		// Protected routes
		//auth.GET("/getallprojects", middleware.AuthMiddleware(), controllers.GetAllProjects)
		//auth.GET("/getprojectbyid/:projectId", middleware.AuthMiddleware(), controllers.GetProjectById)
		auth.GET("/getalltasksbyprojectid/:projectId/:page/:pageSize", middleware.AuthMiddleware(), controllers.GetAllTasksByProjectId)
		auth.POST("/createtask", middleware.AuthMiddleware(), controllers.CreateTask)
	}
}
