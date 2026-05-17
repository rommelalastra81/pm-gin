package controllers

import (
	"net/http"
	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create task with project and user relationships
	task := models.Tasks{
		Title:       req.Title,
		Description: req.Description,
		TaskType:    req.TaskType,
		Status:      req.Status,
		Priority:    req.Priority,
		StartDate:   req.StartDate.Time,
		DueDate:     req.DueDate.Time,
		AssignedTo:  req.AssignedTo, // FK → Users.ID
		ProjectId:   req.ProjectID,  // FK → Projects.ID
	}

	// Only set CompletionDate if it was provided (non-null)
	if req.CompletionDate != nil {
		task.CompletionDate = &req.CompletionDate.Time
	}

	if result := config.DB.Create(&task); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, task)
}
