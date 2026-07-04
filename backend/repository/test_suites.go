package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type testSuiteRepository struct {
	db *gorm.DB
}

func NewTestSuiteRepo(db *gorm.DB) TestSuiteRepository {
	return &testSuiteRepository{db: db}
}

func (r *testSuiteRepository) GetTestSuitesByProject(projectID string) ([]entities.TestSuite, error) {
	var suites []entities.TestSuite
	err := r.db.Where("project_id = ?", projectID).Find(&suites).Error
	return suites, err
}

func (r *testSuiteRepository) GetTestSuiteByID(id string) (entities.TestSuite, error) {
	var suite entities.TestSuite
	err := r.db.First(&suite, "id = ?", id).Error
	return suite, err
}

func (r *testSuiteRepository) CreateTestSuite(projectID string, req entities.CreateTestSuiteRequest) (entities.TestSuite, error) {
	suite := entities.TestSuite{
		ProjectID:   projectID,
		ParentID:    req.ParentID,
		Name:        req.Name,
		Description: req.Description,
	}
	err := r.db.Create(&suite).Error
	return suite, err
}

func (r *testSuiteRepository) UpdateTestSuite(id string, req entities.CreateTestSuiteRequest) (entities.TestSuite, error) {
	var suite entities.TestSuite
	if err := r.db.First(&suite, "id = ?", id).Error; err != nil {
		return suite, err
	}
	suite.ParentID = req.ParentID
	suite.Name = req.Name
	suite.Description = req.Description
	err := r.db.Save(&suite).Error
	return suite, err
}

func (r *testSuiteRepository) DeleteTestSuite(id string) error {
	return r.db.Delete(&entities.TestSuite{}, "id = ?", id).Error
}
