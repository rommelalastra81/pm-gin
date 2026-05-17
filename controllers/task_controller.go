package controllers

import (
	"net/http"
	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"
	"strconv"

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

	c.JSON(http.StatusOK, "ok")
}

func GetAllTasksByProjectId(c *gin.Context) {

	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil || projectId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	page, err := strconv.Atoi(c.Param("page"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(c.Param("pageSize"))
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	offset := (page - 1) * pageSize

	var totalCount int64
	if result := config.DB.Model(&models.Tasks{}).Where("project_id = ?", projectId).Count(&totalCount); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var _tasks []models.Tasks
	if result := config.DB.Where("project_id = ?", projectId).Offset(offset).Limit(pageSize).Find(&_tasks); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var data []dto.TasksResponse
	for _, task := range _tasks {
		var completionDate *dto.DateOnlyTask
		if task.CompletionDate != nil {
			completionDate = &dto.DateOnlyTask{Time: *task.CompletionDate}
		}
		data = append(data, dto.TasksResponse{
			Id:             task.Id,
			Title:          task.Title,
			Description:    task.Description,
			TaskType:       task.TaskType,
			Priority:       task.Priority,
			Status:         task.Status,
			AssignedTo:     task.AssignedTo,
			ProjectID:      task.ProjectId,
			StartDate:      dto.DateOnlyTask{Time: task.StartDate},
			CompletionDate: completionDate,
			DueDate:        dto.DateOnlyTask{Time: task.DueDate},
		})
	}

	totalPages := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	c.JSON(http.StatusOK, gin.H{
		"items":        data,
		"currentPage":  page,
		"pageSize":     pageSize,
		"totalRecords": totalCount,
		"totalPages":   totalPages,
	})

}
