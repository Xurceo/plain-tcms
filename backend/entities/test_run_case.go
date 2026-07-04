package entities

type TestRunCase struct {
	RunID      string `gorm:"primaryKey;column:run_id;type:uuid" json:"run_id"`
	TestCaseID string `gorm:"primaryKey;column:test_case_id;type:uuid" json:"test_case_id"`
}

func (TestRunCase) TableName() string { return "test_run_cases" }
