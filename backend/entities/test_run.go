package entities

import "time"

type TestRun struct {
	ID           string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ProjectID    string     `gorm:"column:project_id;type:uuid;not null" json:"project_id"`
	PlanID       *string    `gorm:"column:plan_id;type:uuid" json:"plan_id"`
	Name         string     `gorm:"not null" json:"name"`
	Status       string     `gorm:"default:pending" json:"status"`
	Environment  *string    `json:"environment"`
	BuildVersion *string    `gorm:"column:build_version" json:"build_version"`
	CreatedBy    *string    `gorm:"column:created_by;type:uuid" json:"created_by"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	CompletedAt  *time.Time `gorm:"column:completed_at" json:"completed_at"`
}

type CreateTestRunRequest struct {
	PlanID       *string `json:"plan_id"`
	Name         string  `json:"name"`
	Environment  *string `json:"environment"`
	BuildVersion *string `json:"build_version"`
}

func (TestRun) TableName() string { return "test_runs" }
