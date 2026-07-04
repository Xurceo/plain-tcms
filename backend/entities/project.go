package entities

import "time"

type Project struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	OrgID       string    `gorm:"column:org_id;type:uuid;not null" json:"org_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description *string   `json:"description"`
	CreatedBy   *string   `gorm:"column:created_by;type:uuid" json:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateProjectRequest struct {
	OrgID       string  `json:"org_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedBy   *string `json:"created_by"`
}

func (Project) TableName() string { return "projects" }
