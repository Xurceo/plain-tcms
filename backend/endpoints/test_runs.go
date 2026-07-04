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

type TestRunHandler struct {
	repo repository.TestRunRepository
}

func NewTestRunHandler(repo repository.TestRunRepository) *TestRunHandler {
	return &TestRunHandler{repo: repo}
}

// GetTestRunsByProject godoc
// @Summary Get test runs by project
// @Tags Test Runs
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {array} entities.TestRun
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-runs [get]
func (h *TestRunHandler) GetTestRunsByProject(c *gin.Context) {
	projectID := c.Param("id")
	runs, err := h.repo.GetTestRunsByProject(projectID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, runs)
}

// GetTestRunByID godoc
// @Summary Get test run by ID
// @Tags Test Runs
// @Produce json
// @Param id path string true "Test Run ID"
// @Success 200 {object} entities.TestRun
// @Failure 404 {object} entities.ErrorResponse
// @Router /test-runs/{id} [get]
func (h *TestRunHandler) GetTestRunByID(c *gin.Context) {
	handleGetByID(c, h.repo.GetTestRunByID)
}

// CreateTestRun godoc
// @Summary Create test run
// @Tags Test Runs
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param test_run body entities.CreateTestRunRequest true "Test Run"
// @Success 201 {object} entities.TestRun
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-runs [post]
func (h *TestRunHandler) CreateTestRun(c *gin.Context) {
	projectID := c.Param("id")
	var req entities.CreateTestRunRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	run, err := h.repo.CreateTestRun(projectID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, run)
}

// UpdateTestRun godoc
// @Summary Update test run
// @Tags Test Runs
// @Accept json
// @Produce json
// @Param id path string true "Test Run ID"
// @Param test_run body entities.CreateTestRunRequest true "Test Run"
// @Success 200 {object} entities.TestRun
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-runs/{id} [put]
func (h *TestRunHandler) UpdateTestRun(c *gin.Context) {
	handleUpdate(c, h.repo.UpdateTestRun)
}

// DeleteTestRun godoc
// @Summary Delete test run
// @Tags Test Runs
// @Param id path string true "Test Run ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-runs/{id} [delete]
func (h *TestRunHandler) DeleteTestRun(c *gin.Context) {
	handleDelete(c, h.repo.DeleteTestRun)
}

// GetCasesByRun godoc
// @Summary Get test cases in a run
// @Tags Test Runs
// @Produce json
// @Param id path string true "Test Run ID"
// @Success 200 {array} entities.TestCase
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-runs/{id}/cases [get]
func (h *TestRunHandler) GetCasesByRun(c *gin.Context) {
	runID := c.Param("id")
	cases, err := h.repo.GetCasesByRun(runID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cases)
}

// AddCaseToRun godoc
// @Summary Add test case to run
// @Tags Test Runs
// @Accept json
// @Produce json
// @Param id path string true "Test Run ID"
// @Param case body object true "Test case ID" SchemaProps: {"test_case_id": "string"}
// @Success 201
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-runs/{id}/cases [post]
func (h *TestRunHandler) AddCaseToRun(c *gin.Context) {
	runID := c.Param("id")
	var body struct {
		TestCaseID string `json:"test_case_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.AddCaseToRun(runID, body.TestCaseID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(201)
}

// RemoveCaseFromRun godoc
// @Summary Remove test case from run
// @Tags Test Runs
// @Param id path string true "Test Run ID"
// @Param test_case_id path string true "Test Case ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-runs/{id}/cases/{test_case_id} [delete]
func (h *TestRunHandler) RemoveCaseFromRun(c *gin.Context) {
	runID := c.Param("id")
	testCaseID := c.Param("test_case_id")
	if err := h.repo.RemoveCaseFromRun(runID, testCaseID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}

// GetResultsByRun godoc
// @Summary Get test results for a run
// @Tags Test Runs
// @Produce json
// @Param id path string true "Test Run ID"
// @Success 200 {array} entities.TestResult
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-runs/{id}/results [get]
func (h *TestRunHandler) GetResultsByRun(c *gin.Context) {
	runID := c.Param("id")
	results, err := h.repo.GetResultsByRun(runID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, results)
}

// AddResultToRun godoc
// @Summary Add test result to run
// @Tags Test Runs
// @Accept json
// @Produce json
// @Param id path string true "Test Run ID"
// @Param result body entities.CreateTestResultRequest true "Test Result"
// @Success 201 {object} entities.TestResult
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-runs/{id}/results [post]
func (h *TestRunHandler) AddResultToRun(c *gin.Context) {
	runID := c.Param("id")
	var req entities.CreateTestResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := h.repo.AddResultToRun(runID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, result)
}
