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
