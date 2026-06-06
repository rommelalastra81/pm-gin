package models

type MemberRoles struct {
	Id uint `json:"id" gorm:"primaryKey"`

	// Many-to-One (User)
	ProjectMemberId uint           `json:"project_member_id"`
	ProjectMembers  ProjectMembers `json:"-" gorm:"foreignKey:ProjectMemberId"` // equivalent to @JsonIgnore

	// Many-to-One (Project)
	RoleId uint  `json:"role_id"`
	Roles  Roles `json:"-" gorm:"foreignKey:RoleId"`
}
