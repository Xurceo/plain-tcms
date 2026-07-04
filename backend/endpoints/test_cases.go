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

package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/entities"
	"github.com/xurceo/plain-tcms/repository"
)

// GetTestCases godoc
// @Summary Get test cases by project id
// @Tags Test Cases
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {array} entities.TestCase
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-cases [get]
func GetTestCases(c *gin.Context) {
	projectID := c.Param("id")
	testCases, err := repository.GetTestCases(projectID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, testCases)
}

// CreateTestCase godoc
// @Summary Create test case
// @Tags Test Cases
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param test_case body entities.CreateTestCaseRequest true "Test Case"
// @Success 201 {object} entities.TestCase
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-cases [post]
func CreateTestCase(c *gin.Context) {
	projectID := c.Param("id")

	var req entities.CreateTestCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tc, err := repository.CreateTestCase(projectID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, tc)
}
