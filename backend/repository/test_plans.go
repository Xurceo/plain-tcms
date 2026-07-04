/*
 * // TCMS - Test Case Management System
 * // Copyright (C) 2026 Pavlo Shnal
 * //
 * // This program is free software: you can redistribute it and/or modify
 * // it under the terms of the GNU Affero General Public License as published
 * // by the Free Software Foundation, either version 3 of the License, or
 * // (at your option) any later version.
 * //
 * // This program is distributed in the hope that it will be useful,
 * // but WITHOUT ANY WARRANTY; without even the implied warranty of
 * // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * // GNU Affero General Public License for more details.
 * //
 * // You should have received a copy of the GNU Affero General Public License
 * // along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

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
