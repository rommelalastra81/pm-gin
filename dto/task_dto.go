package dto

import (
	"fmt"
	"time"
)

// DateOnly is a time.Time that marshals/unmarshals JSON as "YYYY-MM-DD".
type DateOnlyTask struct {
	time.Time
}

const dateLayoutTask = "2006-01-02"

func (d *DateOnlyTask) UnmarshalJSON(b []byte) error {
	s := string(b)
	// handle JSON null — leave time as zero value
	if s == "null" {
		return nil
	}
	// strip surrounding quotes
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	t, err := time.Parse(dateLayoutTask, s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}
	d.Time = t
	return nil
}

func (d DateOnlyTask) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(dateLayoutTask) + `"`), nil
}

type TasksResponse struct {
	Id             uint          `json:"id" gorm:"primaryKey"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	TaskType       string        `json:"task_type"`
	Status         string        `json:"status"`
	Priority       string        `json:"priority"`
	PercentageDone uint          `json:"percentage_done"`
	AssignedTo     uint          `json:"assigned_to"`
	ProjectID      uint          `json:"project_id"`
	StartDate      DateOnlyTask  `json:"start_date" gorm:"type:date"`
	CompletionDate *DateOnlyTask `json:"completion_date" gorm:"type:date"`
	DueDate        DateOnlyTask  `json:"due_date" gorm:"type:date"`
}

type CreateTaskRequest struct {
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	TaskType       string        `json:"task_type"`
	Status         string        `json:"status"`
	Priority       string        `json:"priority"`
	PercentageDone uint          `json:"percentage_done"`
	AssignedTo     uint          `json:"assigned_to"`
	ProjectID      uint          `json:"project_id"`
	StartDate      DateOnlyTask  `json:"start_date" gorm:"type:date"`
	CompletionDate *DateOnlyTask `json:"completion_date" gorm:"type:date"`
	DueDate        DateOnlyTask  `json:"due_date" gorm:"type:date"`
}

type UpdateTaskRequest struct {
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	TaskType       string        `json:"task_type"`
	Status         string        `json:"status"`
	Priority       string        `json:"priority"`
	PercentageDone uint          `json:"percentage_done"`
	AssignedTo     uint          `json:"assigned_to"`
	ProjectID      uint          `json:"project_id"`
	StartDate      DateOnlyTask  `json:"start_date" gorm:"type:date"`
	CompletionDate *DateOnlyTask `json:"completion_date" gorm:"type:date"`
	DueDate        DateOnlyTask  `json:"due_date" gorm:"type:date"`
}

type UpdateTaskStatusRequest struct {
	Status         string        `json:"status" binding:"required"`
	PercentageDone uint          `json:"percentage_done"`
	CompletionDate *DateOnlyTask `json:"completion_date" gorm:"type:date"`
}
