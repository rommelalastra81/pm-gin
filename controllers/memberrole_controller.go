package controllers

import (
	"net/http"
	"pm-gin/config"
	"pm-gin/dto"
	"pm-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateMemberRole(c *gin.Context) {
	projectMemberIdStr := c.Param("projectMemberId")
	projectMemberId, err := strconv.Atoi(projectMemberIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project member ID"})
		return
	}

	var req dto.UpdateMemberRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find member role by project_member_id (matches Java's findByProjectMember_Id)
	var memberRole models.MemberRoles
	if result := config.DB.Where("project_member_id = ?", projectMemberId).First(&memberRole); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member role not found"})
		return
	}

	// Validate project member exists
	var member models.ProjectMembers
	if result := config.DB.First(&member, projectMemberId); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}
	memberRole.ProjectMemberId = uint(projectMemberId)

	// Only update role if provided (matches Java's null check)
	if req.RoleId != 0 {
		var role models.Roles
		if result := config.DB.First(&role, req.RoleId); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}
		memberRole.RoleId = req.RoleId
	}

	if result := config.DB.Save(&memberRole); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member role"})
		return
	}

	// Return DTO matching Java's UpdateMemberRoleDTO
	c.JSON(http.StatusOK, gin.H{
		"id":                memberRole.Id,
		"project_member_id": uint(projectMemberId),
		"role_id":           req.RoleId,
	})
}
