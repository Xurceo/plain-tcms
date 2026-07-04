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

type TestSuiteHandler struct {
	repo repository.TestSuiteRepository
}

func NewTestSuiteHandler(repo repository.TestSuiteRepository) *TestSuiteHandler {
	return &TestSuiteHandler{repo: repo}
}

// GetTestSuitesByProject godoc
// @Summary Get test suites by project
// @Tags Test Suites
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {array} entities.TestSuite
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-suites [get]
func (h *TestSuiteHandler) GetTestSuitesByProject(c *gin.Context) {
	projectID := c.Param("id")
	suites, err := h.repo.GetTestSuitesByProject(projectID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, suites)
}

// GetTestSuiteByID godoc
// @Summary Get test suite by ID
// @Tags Test Suites
// @Produce json
// @Param id path string true "Test Suite ID"
// @Success 200 {object} entities.TestSuite
// @Failure 404 {object} entities.ErrorResponse
// @Router /test-suites/{id} [get]
func (h *TestSuiteHandler) GetTestSuiteByID(c *gin.Context) {
	handleGetByID(c, h.repo.GetTestSuiteByID)
}

// CreateTestSuite godoc
// @Summary Create test suite
// @Tags Test Suites
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param test_suite body entities.CreateTestSuiteRequest true "Test Suite"
// @Success 201 {object} entities.TestSuite
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-suites [post]
func (h *TestSuiteHandler) CreateTestSuite(c *gin.Context) {
	projectID := c.Param("id")
	var req entities.CreateTestSuiteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	suite, err := h.repo.CreateTestSuite(projectID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, suite)
}

// UpdateTestSuite godoc
// @Summary Update test suite
// @Tags Test Suites
// @Accept json
// @Produce json
// @Param id path string true "Test Suite ID"
// @Param test_suite body entities.CreateTestSuiteRequest true "Test Suite"
// @Success 200 {object} entities.TestSuite
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-suites/{id} [put]
func (h *TestSuiteHandler) UpdateTestSuite(c *gin.Context) {
	handleUpdate(c, h.repo.UpdateTestSuite)
}

// DeleteTestSuite godoc
// @Summary Delete test suite
// @Tags Test Suites
// @Param id path string true "Test Suite ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-suites/{id} [delete]
func (h *TestSuiteHandler) DeleteTestSuite(c *gin.Context) {
	handleDelete(c, h.repo.DeleteTestSuite)
}
