package dto

import (
	"time"
)

type TaskCommentResponse struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	TaskId    uint      `json:"task_id"`
	FullName  string    `json:"full_name"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTaskCommentRequest struct {
	UserId  uint   `json:"user_id"`
	TaskId  uint   `json:"task_id"`
	Comment string `json:"comment"`
}
