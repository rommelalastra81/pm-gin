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
		auth.GET("/getpaginatedtasksbyprojectid/:projectId/:page/:pageSize", middleware.AuthMiddleware(), controllers.GetAllTasksByProjectId)
		auth.GET("/getpaginatedtasksbyowner/:projectId/:assignedTo/:page/:pageSize", middleware.AuthMiddleware(), controllers.GetAllTasksByOwner)
		auth.POST("/createtask", middleware.AuthMiddleware(), controllers.CreateTask)
		auth.PUT("/updatetask/:taskId", middleware.AuthMiddleware(), controllers.UpdateTask)
		auth.GET("/gettaskbyid/:taskId", middleware.AuthMiddleware(), controllers.GetTaskById)
		auth.PATCH("/updatetaskstatus/:taskId", middleware.AuthMiddleware(), controllers.UpdateTaskStatus)
	}
}
