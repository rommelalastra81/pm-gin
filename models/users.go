package models

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	JobRole  string `json:"job_role"`

	// One-to-Many
	ProjectMembers []ProjectMembers `json:"project_members" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Tasks          []Tasks          `json:"tasks" gorm:"foreignKey:AssignedTo;constraint:OnDelete:CASCADE"`
}
