package controllers

import (
	"net/http"
	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {

	var req dto.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create project
	project := models.Projects{
		Name:             req.Name,
		Description:      req.Description,
		Status:           req.Status,
		StartDate:        req.StartDate.Time,
		TargetCompletion: req.TargetCompletion.Time,
	}

	if result := config.DB.Create(&project); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// GetAllProjects returns an array of projects
// GET /api/Project/getallprojects
func GetAllProjects(c *gin.Context) {

	var _projects []models.Projects

	if result := config.DB.Find(&_projects); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	var response []dto.ProjectsResponse

	for _, project := range _projects {

		response = append(response, dto.ProjectsResponse{
			Name:             project.Name,
			Description:      project.Description,
			Status:           project.Status,
			StartDate:        dto.DateOnly{Time: project.StartDate},
			TargetCompletion: dto.DateOnly{Time: project.TargetCompletion},
		})
	}

	c.JSON(http.StatusOK, response)

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
		StartDate:        dto.DateOnly{Time: _projects.StartDate},
		TargetCompletion: dto.DateOnly{Time: _projects.TargetCompletion},
	})
}
