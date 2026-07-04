package entities

import "time"

type TestSuite struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ProjectID   string    `gorm:"column:project_id;type:uuid;not null" json:"project_id"`
	ParentID    *string   `gorm:"column:parent_id;type:uuid" json:"parent_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateTestSuiteRequest struct {
	ParentID    *string `json:"parent_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (TestSuite) TableName() string { return "test_suites" }
