package entities

import "time"

type TestResult struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	RunID      string    `gorm:"column:run_id;type:uuid;not null" json:"run_id"`
	TestCaseID string    `gorm:"column:test_case_id;type:uuid;not null" json:"test_case_id"`
	Status     string    `gorm:"default:untested" json:"status"`
	Comment    *string   `json:"comment"`
	ExecutedBy *string   `gorm:"column:executed_by;type:uuid" json:"executed_by"`
	DurationMs *int      `gorm:"column:duration_ms" json:"duration_ms"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateTestResultRequest struct {
	TestCaseID string  `json:"test_case_id"`
	Status     string  `json:"status"`
	Comment    *string `json:"comment"`
	DurationMs *int    `json:"duration_ms"`
}

func (TestResult) TableName() string { return "test_results" }
