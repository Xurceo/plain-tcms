package entities

type OrganizationMember struct {
	OrgID  string `gorm:"primaryKey;column:org_id;type:uuid" json:"org_id"`
	UserID string `gorm:"primaryKey;column:user_id;type:uuid" json:"user_id"`
	Role   string `gorm:"not null" json:"role"`
}

type CreateOrganizationMemberRequest struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

func (OrganizationMember) TableName() string { return "organization_members" }
