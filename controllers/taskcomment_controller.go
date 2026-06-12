package controllers

import (
	"net/http"
	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTaskCommentsByTaskId(c *gin.Context) {

	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil || taskId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
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
	if result := config.DB.Model(&models.TaskComments{}).Where("task_id = ?", taskId).Count(&totalCount); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var _tasksComments []models.TaskComments
	if result := config.DB.Preload("Users").Where("task_id = ?", taskId).Offset(offset).Limit(pageSize).Find(&_tasksComments); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	totalPages := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	var data []dto.TaskCommentResponse
	for _, taskComment := range _tasksComments {

		data = append(data, dto.TaskCommentResponse{
			Id:        taskComment.Id,
			UserId:    taskComment.UserId,
			TaskId:    taskComment.TaskId,
			FullName:  taskComment.Users.FullName,
			Comment:   taskComment.Comment,
			CreatedAt: taskComment.CreatedAt,
			UpdatedAt: taskComment.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"items":        data,
		"currentPage":  page,
		"pageSize":     pageSize,
		"totalRecords": totalCount,
		"totalPages":   totalPages,
	})
}
