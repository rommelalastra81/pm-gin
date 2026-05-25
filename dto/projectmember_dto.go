package dto

type AddProjectMemberRequest struct {
	UserId uint `json:"user_id"`
	RoleId uint `json:"role_id"`
}

type ProjectMemberResponse struct {
	Id        uint    `json:"id"`
	UserId    uint    `json:"user_id"`
	FullName  *string `json:"full_name"`
	ProjectId uint    `json:"project_id"`
	RoleId    *uint   `json:"role_id"`
	Role      *string `json:"role"`
}

type UserProjectMemberDTO struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	JobRole  string `json:"job_role"`
}
