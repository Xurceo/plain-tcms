package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type testCaseRepository struct {
	db *gorm.DB
}

func NewTestCaseRepo(db *gorm.DB) TestCaseRepository {
	return &testCaseRepository{db: db}
}

func (r *testCaseRepository) GetTestCases(projectID string) ([]entities.TestCase, error) {
	var testCases []entities.TestCase
	err := r.db.Where("project_id = ?", projectID).Find(&testCases).Error
	return testCases, err
}

func (r *testCaseRepository) GetTestCaseByID(id string) (entities.TestCase, error) {
	var tc entities.TestCase
	err := r.db.First(&tc, "id = ?", id).Error
	return tc, err
}

func (r *testCaseRepository) CreateTestCase(projectID string, req entities.CreateTestCaseRequest) (entities.TestCase, error) {
	tc := entities.TestCase{
		ProjectID:     projectID,
		SuiteID:       req.SuiteID,
		Title:         req.Title,
		Description:   req.Description,
		Preconditions: req.Preconditions,
		Expected:      req.Expected,
		Priority:      req.Priority,
		Type:          req.Type,
		Tags:          req.Tags,
	}
	err := r.db.Create(&tc).Error
	return tc, err
}

func (r *testCaseRepository) UpdateTestCase(id string, req entities.CreateTestCaseRequest) (entities.TestCase, error) {
	var tc entities.TestCase
	if err := r.db.First(&tc, "id = ?", id).Error; err != nil {
		return tc, err
	}
	tc.SuiteID = req.SuiteID
	tc.Title = req.Title
	tc.Description = req.Description
	tc.Preconditions = req.Preconditions
	tc.Expected = req.Expected
	tc.Priority = req.Priority
	tc.Type = req.Type
	tc.Tags = req.Tags
	err := r.db.Save(&tc).Error
	return tc, err
}

func (r *testCaseRepository) DeleteTestCase(id string) error {
	return r.db.Delete(&entities.TestCase{}, "id = ?", id).Error
}

func (r *testCaseRepository) GetTestCaseHistory(testCaseID string) ([]entities.TestCaseHistory, error) {
	var history []entities.TestCaseHistory
	err := r.db.Where("test_case_id = ?", testCaseID).Order("changed_at DESC").Find(&history).Error
	return history, err
}
