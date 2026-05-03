package models

type Roles struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Role string `json:"role"`

	// One-to-Many
	MemberRoles []MemberRoles `json:"tasks" gorm:"foreignKey:RoleId;constraint:OnDelete:CASCADE"`
}
