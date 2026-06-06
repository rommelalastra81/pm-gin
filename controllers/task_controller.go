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
		Title:          req.Title,
		Description:    req.Description,
		TaskType:       req.TaskType,
		Status:         req.Status,
		Priority:       req.Priority,
		PercentageDone: req.PercentageDone,
		StartDate:      req.StartDate.Time,
		DueDate:        req.DueDate.Time,
		AssignedTo:     req.AssignedTo, // FK → Users.ID
		ProjectId:      req.ProjectID,  // FK → Projects.ID
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

		var _user []models.Users
		if resultUser := config.DB.First(&_user, task.AssignedTo); resultUser.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": resultUser.Error.Error()})
			return
		}

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
			PercentageDone: task.PercentageDone,
			AssignedTo:     task.AssignedTo,
			ProjectID:      task.ProjectId,
			StartDate:      dto.DateOnlyTask{Time: task.StartDate},
			CompletionDate: completionDate,
			DueDate:        dto.DateOnlyTask{Time: task.DueDate},
			UserId:         task.AssignedTo,
			FullName:       _user[0].FullName,
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

func GetAllTasksByOwner(c *gin.Context) {

	projectId, err := strconv.Atoi(c.Param("projectId"))
	assignedTo, err := strconv.Atoi(c.Param("assignedTo"))
	if err != nil || projectId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	if err != nil || assignedTo < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assigned-to ID"})
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
	if result := config.DB.Model(&models.Tasks{}).Where("project_id = ? AND assigned_to = ?", projectId, assignedTo).Count(&totalCount); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var _tasks []models.Tasks
	if result := config.DB.Where("project_id = ? AND assigned_to = ?", projectId, assignedTo).Offset(offset).Limit(pageSize).Find(&_tasks); result.Error != nil {
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
			PercentageDone: task.PercentageDone,
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

func UpdateTask(c *gin.Context) {
	taskIdStr := c.Param("taskId")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Tasks
	if result := config.DB.First(&task, taskId); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Update task fields
	task.Title = req.Title
	task.Description = req.Description
	task.TaskType = req.TaskType
	task.Status = req.Status
	task.Priority = req.Priority
	task.PercentageDone = req.PercentageDone
	task.StartDate = req.StartDate.Time
	task.DueDate = req.DueDate.Time
	task.AssignedTo = req.AssignedTo
	task.ProjectId = req.ProjectID

	// Update CompletionDate if provided, or set to null
	if req.CompletionDate != nil {
		task.CompletionDate = &req.CompletionDate.Time
	} else {
		task.CompletionDate = nil
	}

	if result := config.DB.Save(&task); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func UpdateTaskStatus(c *gin.Context) {
	taskIdStr := c.Param("taskId")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var req dto.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Tasks
	if result := config.DB.First(&task, taskId); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Status = req.Status
	task.PercentageDone = req.PercentageDone

	// Update CompletionDate if provided, or set to null
	if req.CompletionDate != nil {
		task.CompletionDate = &req.CompletionDate.Time
	} else {
		task.CompletionDate = nil
	}

	if result := config.DB.Save(&task); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task status updated successfully"})
}

func GetTaskById(c *gin.Context) {

	taskIdStr := c.Param("taskId")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Tasks
	if result := config.DB.First(&task, taskId); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var completionDate *dto.DateOnlyTask
	if task.CompletionDate != nil {
		completionDate = &dto.DateOnlyTask{Time: *task.CompletionDate}
	}

	c.JSON(http.StatusOK, dto.TasksResponse{
		Id:             task.Id,
		Title:          task.Title,
		Description:    task.Description,
		TaskType:       task.TaskType,
		Priority:       task.Priority,
		Status:         task.Status,
		PercentageDone: task.PercentageDone,
		AssignedTo:     task.AssignedTo,
		ProjectID:      task.ProjectId,
		StartDate:      dto.DateOnlyTask{Time: task.StartDate},
		CompletionDate: completionDate,
		DueDate:        dto.DateOnlyTask{Time: task.DueDate},
	})
}
