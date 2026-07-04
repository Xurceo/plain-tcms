package entities

import (
	"time"

	"gorm.io/datatypes"
)

type TestCaseHistory struct {
	ID         string            `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TestCaseID string            `gorm:"column:test_case_id;type:uuid;not null" json:"test_case_id"`
	Snapshot   datatypes.JSON    `gorm:"type:jsonb;not null" json:"snapshot"`
	ChangedBy  *string           `gorm:"column:changed_by;type:uuid" json:"changed_by"`
	ChangedAt  time.Time         `gorm:"autoCreateTime" json:"changed_at"`
}

func (TestCaseHistory) TableName() string { return "test_case_history" }
