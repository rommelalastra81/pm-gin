package controllers

import (
	"net/http"
	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTimeEntries(c *gin.Context) {

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
	if result := config.DB.Model(&models.TimeEntries{}).Where("task_id = ?", taskId).Count(&totalCount); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var _timeEntries []models.TimeEntries
	if result := config.DB.Where("task_id = ?", taskId).Offset(offset).Limit(pageSize).Find(&_timeEntries); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	totalPages := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	data := make([]dto.TimeEntryResponse, 0)
	for _, timeEntry := range _timeEntries {

		data = append(data, dto.TimeEntryResponse{
			Id:         timeEntry.Id,
			UserId:     timeEntry.UserId,
			TaskId:     timeEntry.TaskId,
			Activities: timeEntry.Activities,
			Date:       dto.DateOnlyTimeEntry{Time: timeEntry.Date},
			StartTime:  dto.TimeOnlyTimeEntry{Time: timeEntry.StartTime},
			EndTime:    dto.TimeOnlyTimeEntry{Time: timeEntry.EndTime},
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
