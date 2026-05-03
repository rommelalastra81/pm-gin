package dto

type RegisterUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	JobRole  string `json:"job_role" binding:"required"`
}

type RegisterSuccessful struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	JobRole  string `json:"job_role" binding:"required"`
	Message  string `json:"message"`
}
