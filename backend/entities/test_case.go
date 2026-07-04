package entities

import "time"

type TestCase struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ProjectID     string    `gorm:"column:project_id;type:uuid;not null" json:"project_id"`
	SuiteID       *string   `gorm:"column:suite_id;type:uuid" json:"suite_id"`
	Title         string    `gorm:"not null" json:"title"`
	Description   *string   `json:"description"`
	Preconditions *string   `json:"preconditions"`
	Steps         string    `gorm:"type:jsonb;default:'[]'" json:"steps"`
	Expected      *string   `json:"expected"`
	Status        string    `gorm:"default:draft" json:"status"`
	Priority      string    `gorm:"default:medium" json:"priority"`
	Type          *string   `json:"type"`
	Tags          []string  `gorm:"type:jsonb;serializer:json;default:'[]'" json:"tags"`
	CreatedBy     *string   `gorm:"column:created_by;type:uuid" json:"created_by"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type CreateTestCaseRequest struct {
	SuiteID       *string  `json:"suite_id"`
	Title         string   `json:"title"`
	Description   *string  `json:"description"`
	Preconditions *string  `json:"preconditions"`
	Expected      *string  `json:"expected"`
	Priority      string   `json:"priority"`
	Type          *string  `json:"type"`
	Tags          []string `json:"tags"`
}

func (TestCase) TableName() string { return "test_cases" }
