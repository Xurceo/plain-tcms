package repository

import (
	"github.com/xurceo/plain-tcms/db"
	"github.com/xurceo/plain-tcms/entities"
)

func GetTestCases(projectID string) ([]entities.TestCase, error) {
	var testCases []entities.TestCase
	err := db.DB.Where("project_id = ?", projectID).Find(&testCases).Error
	return testCases, err
}

func CreateTestCase(projectID string, req entities.CreateTestCaseRequest) (entities.TestCase, error) {
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
	err := db.DB.Create(&tc).Error
	return tc, err
}
