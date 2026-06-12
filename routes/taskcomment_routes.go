package routes

import (
	"pm-gin/controllers"
	"pm-gin/middleware"

	"github.com/gin-gonic/gin"
)

func TaskCommentRoutes(router *gin.Engine) {
	taskComments := router.Group("/api/TaskComment")
	{

		taskComments.GET("/gettaskcommentsbytaskid/:taskId/:page/:pageSize", middleware.AuthMiddleware(), controllers.GetTaskCommentsByTaskId)
		//taskComments.POST("/createtask", middleware.AuthMiddleware(), controllers.CreateTask)

	}
}
