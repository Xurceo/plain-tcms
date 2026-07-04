package entities

import "time"

type TestPlan struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ProjectID   string    `gorm:"column:project_id;type:uuid;not null" json:"project_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateTestPlanRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (TestPlan) TableName() string { return "test_plans" }
