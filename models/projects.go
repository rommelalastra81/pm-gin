package models

import "time"

type Projects struct {
	Id               uint      `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	StartDate        time.Time `json:"start_date" gorm:"type:date"`
	TargetCompletion time.Time `json:"target_completion" gorm:"type:date"`

	// One-to-Many
	ProjectMembers []ProjectMembers `json:"project_members" gorm:"foreignKey:ProjectId;constraint:OnDelete:CASCADE"`
	Tasks          []Tasks          `json:"tasks" gorm:"foreignKey:ProjectId;constraint:OnDelete:CASCADE"`
}
