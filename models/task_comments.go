package models

import "time"

type TaskComments struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Many-to-One (Users)
	UserId uint  `json:"user_id"`
	Users  Users `json:"-" gorm:"foreignKey:UserId"` // equivalent to @JsonIgnore

	// Many-to-One (Tasks)
	TaskId uint  `json:"task_id"`
	Tasks  Tasks `json:"-" gorm:"foreignKey:TaskId"`
}
