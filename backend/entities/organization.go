package entities

import "time"

type Organization struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateOrganizationRequest struct {
	Name string `json:"name"`
}

func (Organization) TableName() string { return "organizations" }
