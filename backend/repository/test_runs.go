package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type testRunRepository struct {
	db *gorm.DB
}

func NewTestRunRepo(db *gorm.DB) TestRunRepository {
	return &testRunRepository{db: db}
}

func (r *testRunRepository) GetTestRunsByProject(projectID string) ([]entities.TestRun, error) {
	var runs []entities.TestRun
	err := r.db.Where("project_id = ?", projectID).Find(&runs).Error
	return runs, err
}

func (r *testRunRepository) GetTestRunByID(id string) (entities.TestRun, error) {
	var run entities.TestRun
	err := r.db.First(&run, "id = ?", id).Error
	return run, err
}

func (r *testRunRepository) CreateTestRun(projectID string, req entities.CreateTestRunRequest) (entities.TestRun, error) {
	run := entities.TestRun{
		ProjectID:    projectID,
		PlanID:       req.PlanID,
		Name:         req.Name,
		Environment:  req.Environment,
		BuildVersion: req.BuildVersion,
	}
	err := r.db.Create(&run).Error
	return run, err
}

func (r *testRunRepository) UpdateTestRun(id string, req entities.CreateTestRunRequest) (entities.TestRun, error) {
	var run entities.TestRun
	if err := r.db.First(&run, "id = ?", id).Error; err != nil {
		return run, err
	}
	run.PlanID = req.PlanID
	run.Name = req.Name
	run.Environment = req.Environment
	run.BuildVersion = req.BuildVersion
	err := r.db.Save(&run).Error
	return run, err
}

func (r *testRunRepository) DeleteTestRun(id string) error {
	return r.db.Delete(&entities.TestRun{}, "id = ?", id).Error
}

func (r *testRunRepository) GetCasesByRun(runID string) ([]entities.TestCase, error) {
	var cases []entities.TestCase
	err := r.db.
		Joins("INNER JOIN test_run_cases ON test_run_cases.test_case_id = test_cases.id").
		Where("test_run_cases.run_id = ?", runID).
		Find(&cases).Error
	return cases, err
}

func (r *testRunRepository) AddCaseToRun(runID string, testCaseID string) error {
	runCase := entities.TestRunCase{RunID: runID, TestCaseID: testCaseID}
	return r.db.Create(&runCase).Error
}

func (r *testRunRepository) RemoveCaseFromRun(runID string, testCaseID string) error {
	return r.db.Where("run_id = ? AND test_case_id = ?", runID, testCaseID).Delete(&entities.TestRunCase{}).Error
}

func (r *testRunRepository) GetResultsByRun(runID string) ([]entities.TestResult, error) {
	var results []entities.TestResult
	err := r.db.Where("run_id = ?", runID).Find(&results).Error
	return results, err
}

func (r *testRunRepository) AddResultToRun(runID string, req entities.CreateTestResultRequest) (entities.TestResult, error) {
	result := entities.TestResult{
		RunID:      runID,
		TestCaseID: req.TestCaseID,
		Status:     req.Status,
		Comment:    req.Comment,
		DurationMs: req.DurationMs,
	}
	err := r.db.Create(&result).Error
	return result, err
}
