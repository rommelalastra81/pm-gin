package models

import (
	"time"
)

type Tasks struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	TaskType       string     `json:"task_type"`
	Status         string     `json:"status"`
	Priority       string     `json:"priority"`
	StartDate      time.Time  `json:"start_date" gorm:"type:date"`
	CompletionDate *time.Time `json:"completion_date" gorm:"type:date"`
	DueDate        time.Time  `json:"due_date" gorm:"type:date"`

	// Many-to-One (User)
	AssignedTo uint  `json:"assigned_to"`
	Users      Users `json:"-" gorm:"foreignKey:AssignedTo"` // equivalent to @JsonIgnore

	// Many-to-One (Project)
	ProjectId uint     `json:"project_id"`
	Projects  Projects `json:"-" gorm:"foreignKey:ProjectId"`
}
