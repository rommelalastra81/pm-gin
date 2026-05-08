package dto

import "time"

type ProjectsResponse struct {
	Id               uint      `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	StartDate        time.Time `json:"start_date" gorm:"type:date"`
	TargetCompletion time.Time `json:"target_completion" gorm:"type:date"`
}
