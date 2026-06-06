package models

import "time"

type TimeEntries struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Activities string    `json:"activities"`
	Date       time.Time `json:"date" gorm:"type:date"`
	StartTime  time.Time `json:"start_time" gorm:"type:time"`
	EndTime    time.Time `json:"end_time" gorm:"type:time"`

	// Many-to-One (Users)
	UserId uint  `json:"user_id"`
	Users  Users `json:"-" gorm:"foreignKey:UserId"` // equivalent to @JsonIgnore

	// Many-to-One (Tasks)
	TaskId uint  `json:"task_id"`
	Tasks  Tasks `json:"-" gorm:"foreignKey:TaskId"`
}
