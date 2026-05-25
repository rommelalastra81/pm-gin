package models

type ProjectMembers struct {
	Id uint `json:"id" gorm:"primaryKey"`

	// Many-to-One (User)
	UserId uint  `json:"user_id"`
	Users  Users `json:"-" gorm:"foreignKey:UserId"` // equivalent to @JsonIgnore

	// Many-to-One (Project)
	ProjectId uint     `json:"project_id"`
	Projects  Projects `json:"-" gorm:"foreignKey:ProjectId"`

	// One-to-Many (MemberRoles)
	MemberRoles []MemberRoles `json:"member_roles" gorm:"foreignKey:ProjectMemberId;constraint:OnDelete:CASCADE"`
}
