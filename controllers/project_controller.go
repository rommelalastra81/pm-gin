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

// GetAllProjects returns a paginated list of projects.
// GET /api/Project/getallprojects/:page/:pageSize
func GetAllProjects(c *gin.Context) {

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
	if result := config.DB.Model(&models.Projects{}).Count(&totalCount); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var _projects []models.Projects
	if result := config.DB.Offset(offset).Limit(pageSize).Find(&_projects); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var data []dto.ProjectsResponse
	for _, project := range _projects {
		data = append(data, dto.ProjectsResponse{
			Id:               project.Id,
			Name:             project.Name,
			Description:      project.Description,
			Status:           project.Status,
			StartDate:        dto.DateOnly{Time: project.StartDate},
			TargetCompletion: dto.DateOnly{Time: project.TargetCompletion},
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

// returns non paginated list of all projects
func GetProjects(c *gin.Context) {

	var _projects []models.Projects
	if result := config.DB.Find(&_projects); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var data []dto.ProjectsResponse
	for _, project := range _projects {
		data = append(data, dto.ProjectsResponse{
			Id:               project.Id,
			Name:             project.Name,
			Description:      project.Description,
			Status:           project.Status,
			StartDate:        dto.DateOnly{Time: project.StartDate},
			TargetCompletion: dto.DateOnly{Time: project.TargetCompletion},
		})
	}
	c.JSON(http.StatusOK, data)
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
		Id:               _projects.Id,
		Name:             _projects.Name,
		Description:      _projects.Description,
		Status:           _projects.Status,
		StartDate:        dto.DateOnly{Time: _projects.StartDate},
		TargetCompletion: dto.DateOnly{Time: _projects.TargetCompletion},
	})
}
