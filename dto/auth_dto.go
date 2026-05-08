package dto

// LoginRequest is the payload for user login (email as username).
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	//User    UserResponse `json:"user"`
	Type     string `json:"type"`
	UserId   uint   `json:"user_id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

// UserResponse is a safe user representation (no password).
type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	JobRole  string `json:"job_role"`
}
