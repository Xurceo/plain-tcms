package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type testPlanRepository struct {
	db *gorm.DB
}

func NewTestPlanRepo(db *gorm.DB) TestPlanRepository {
	return &testPlanRepository{db: db}
}

func (r *testPlanRepository) GetTestPlansByProject(projectID string) ([]entities.TestPlan, error) {
	var plans []entities.TestPlan
	err := r.db.Where("project_id = ?", projectID).Find(&plans).Error
	return plans, err
}

func (r *testPlanRepository) GetTestPlanByID(id string) (entities.TestPlan, error) {
	var plan entities.TestPlan
	err := r.db.First(&plan, "id = ?", id).Error
	return plan, err
}

func (r *testPlanRepository) CreateTestPlan(projectID string, req entities.CreateTestPlanRequest) (entities.TestPlan, error) {
	plan := entities.TestPlan{
		ProjectID:   projectID,
		Name:        req.Name,
		Description: req.Description,
	}
	err := r.db.Create(&plan).Error
	return plan, err
}

func (r *testPlanRepository) UpdateTestPlan(id string, req entities.CreateTestPlanRequest) (entities.TestPlan, error) {
	var plan entities.TestPlan
	if err := r.db.First(&plan, "id = ?", id).Error; err != nil {
		return plan, err
	}
	plan.Name = req.Name
	plan.Description = req.Description
	err := r.db.Save(&plan).Error
	return plan, err
}

func (r *testPlanRepository) DeleteTestPlan(id string) error {
	return r.db.Delete(&entities.TestPlan{}, "id = ?", id).Error
}
