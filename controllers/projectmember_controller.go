package controllers

import (
	"fmt"
	"net/http"
	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProjectMembers(c *gin.Context) {

	// Get projectId from URL path
	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil || projectId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Bind JSON array of member requests
	var requests []dto.AddProjectMemberRequest
	if err := c.ShouldBindJSON(&requests); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(requests) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one member request is required"})
		return
	}

	// Validate that the project exists (once, like in Java)
	var project models.Projects
	if result := config.DB.First(&project, projectId); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Project not found with id: %d", projectId)})
		return
	}

	// Use a transaction to create all project members atomically
	tx := config.DB.Begin()

	var savedMembers []models.ProjectMembers

	for _, req := range requests {
		// Validate that the user exists
		var user models.Users
		if result := config.DB.First(&user, req.UserId); result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found with id: %d", req.UserId)})
			return
		}

		// Create project member
		projectMember := models.ProjectMembers{
			UserId:    req.UserId,
			ProjectId: uint(projectId),
		}

		if result := tx.Create(&projectMember); result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add project member"})
			return
		}

		// Handle Role if provided
		if req.RoleId != 0 {
			// Validate that the role exists
			var role models.Roles
			if result := config.DB.First(&role, req.RoleId); result.Error != nil {
				tx.Rollback()
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Role not found with id: %d", req.RoleId)})
				return
			}

			// Create the member role association
			memberRole := models.MemberRoles{
				ProjectMemberId: projectMember.Id,
				RoleId:          req.RoleId,
			}

			if result := tx.Create(&memberRole); result.Error != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign role to project member"})
				return
			}
		}

		savedMembers = append(savedMembers, projectMember)
	}

	tx.Commit()

	// Convert to response DTOs
	data := make([]dto.ProjectMemberResponse, 0)
	for _, member := range savedMembers {
		data = append(data, dto.ProjectMemberResponse{
			Id:        member.Id,
			UserId:    member.UserId,
			ProjectId: member.ProjectId,
		})
	}

	c.JSON(http.StatusOK, data)
}

func GetMembersByProjectId(c *gin.Context) {

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
	if result := config.DB.Model(&models.ProjectMembers{}).Where("project_id = ?", projectId).Count(&totalCount); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var _projectMembers []models.ProjectMembers
	if result := config.DB.
		Preload("Users").
		Preload("MemberRoles").
		Preload("MemberRoles.Roles").
		Where("project_id = ?", projectId).
		Offset(offset).Limit(pageSize).
		Find(&_projectMembers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	data := make([]dto.ProjectMemberResponse, 0)
	for _, member := range _projectMembers {
		resp := dto.ProjectMemberResponse{
			Id:        member.Id,
			UserId:    member.UserId,
			ProjectId: member.ProjectId,
		}

		// Get user full name
		if member.Users.Id != 0 {
			fullName := member.Users.FullName
			resp.FullName = &fullName
		}

		// Get role info from first MemberRole (matching Java logic)
		if len(member.MemberRoles) > 0 {
			mr := member.MemberRoles[0]
			if mr.Roles.Id != 0 {
				roleId := mr.Roles.Id
				roleName := mr.Roles.Role
				resp.RoleId = &roleId
				resp.Role = &roleName
			}
		}

		data = append(data, resp)
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

func GetUsersNotOnProject(c *gin.Context) {

	projectId, err := strconv.Atoi(c.Param("projectId"))
	if err != nil || projectId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Equivalent to:
	// SELECT * FROM users WHERE id NOT IN
	//   (SELECT user_id FROM project_members WHERE project_id = ?)
	var users []models.Users
	if result := config.DB.
		Where("id NOT IN (?)",
			config.DB.Model(&models.ProjectMembers{}).
				Select("user_id").
				Where("project_id = ?", projectId),
		).Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	data := make([]dto.UserProjectMemberDTO, 0)
	for _, user := range users {
		data = append(data, dto.UserProjectMemberDTO{
			Id:       user.Id,
			FullName: user.FullName,
			Email:    user.Email,
			JobRole:  user.JobRole,
		})
	}

	c.JSON(http.StatusOK, data)
}

func GetProjectMemberById(c *gin.Context) {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil || Id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project member ID"})
		return
	}

	var projectMember models.ProjectMembers
	if result := config.DB.
		Preload("Users").
		Preload("MemberRoles").
		Preload("MemberRoles.Roles").
		First(&projectMember, Id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Project member not found with id: %d", Id)})
		return
	}

	resp := dto.ProjectMemberResponse{
		Id:        projectMember.Id,
		UserId:    projectMember.UserId,
		ProjectId: projectMember.ProjectId,
	}

	// Get user full name
	if projectMember.Users.Id != 0 {
		fullName := projectMember.Users.FullName
		resp.FullName = &fullName
	}

	// Get role info from first MemberRole (matching GetMembersByProjectId logic)
	if len(projectMember.MemberRoles) > 0 {
		mr := projectMember.MemberRoles[0]
		if mr.Roles.Id != 0 {
			roleId := mr.Roles.Id
			roleName := mr.Roles.Role
			resp.RoleId = &roleId
			resp.Role = &roleName
		}
	}

	c.JSON(http.StatusOK, resp)
}
