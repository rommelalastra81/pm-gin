package dto

type UpdateMemberRoleRequest struct {
	ProjectMemberId uint `json:"project_member_id"`
	RoleId          uint `json:"role_id"`
}
