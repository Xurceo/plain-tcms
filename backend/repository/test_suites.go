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
