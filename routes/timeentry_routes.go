package routes

import (
	"pm-gin/controllers"
	"pm-gin/middleware"

	"github.com/gin-gonic/gin"
)

func TTimeEntryRoutes(router *gin.Engine) {
	timeEntries := router.Group("/api/TimeEntry")
	{

		timeEntries.GET("/gettimeentries/:taskId/:page/:pageSize", middleware.AuthMiddleware(), controllers.GetTimeEntries)
		//timeEntries.POST("/createtask", middleware.AuthMiddleware(), controllers.CreateTask)

	}
}
