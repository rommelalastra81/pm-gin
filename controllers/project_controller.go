package controllers

import (
	"net/http"
	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProjects(c *gin.Context) {

}

// GetProjectById returns a project by its ID.
// GET /api/Project/getprojectbyid/:projectId
func GetProjectById(c *gin.Context) {
	projectIdStr := c.Param("projectId")
	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var _projects models.Projects
	if result := config.DB.First(&_projects, projectId); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ProjectsResponse{
		Id:               _projects.ID,
		Name:             _projects.Name,
		Description:      _projects.Description,
		Status:           _projects.Status,
		StartDate:        _projects.StartDate,
		TargetCompletion: _projects.TargetCompletion,
	})
}
