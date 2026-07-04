package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/entities"
	"github.com/xurceo/plain-tcms/repository"
)

type TestCaseHandler struct {
	repo repository.TestCaseRepository
}

func NewTestCaseHandler(repo repository.TestCaseRepository) *TestCaseHandler {
	return &TestCaseHandler{repo: repo}
}

// GetTestCases godoc
// @Summary Get test cases by project id
// @Tags Test Cases
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {array} entities.TestCase
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-cases [get]
func (h *TestCaseHandler) GetTestCases(c *gin.Context) {
	projectID := c.Param("id")
	testCases, err := h.repo.GetTestCases(projectID)
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
func (h *TestCaseHandler) CreateTestCase(c *gin.Context) {
	projectID := c.Param("id")

	var req entities.CreateTestCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tc, err := h.repo.CreateTestCase(projectID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, tc)
}

// GetTestCaseByID godoc
// @Summary Get test case by ID
// @Tags Test Cases
// @Produce json
// @Param id path string true "Test Case ID"
// @Success 200 {object} entities.TestCase
// @Failure 404 {object} entities.ErrorResponse
// @Router /test-cases/{id} [get]
func (h *TestCaseHandler) GetTestCaseByID(c *gin.Context) {
	handleGetByID(c, h.repo.GetTestCaseByID)
}

// UpdateTestCase godoc
// @Summary Update test case
// @Tags Test Cases
// @Accept json
// @Produce json
// @Param id path string true "Test Case ID"
// @Param test_case body entities.CreateTestCaseRequest true "Test Case"
// @Success 200 {object} entities.TestCase
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-cases/{id} [put]
func (h *TestCaseHandler) UpdateTestCase(c *gin.Context) {
	handleUpdate(c, h.repo.UpdateTestCase)
}

// DeleteTestCase godoc
// @Summary Delete test case
// @Tags Test Cases
// @Param id path string true "Test Case ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-cases/{id} [delete]
func (h *TestCaseHandler) DeleteTestCase(c *gin.Context) {
	handleDelete(c, h.repo.DeleteTestCase)
}

// GetTestCaseHistory godoc
// @Summary Get test case history
// @Tags Test Cases
// @Produce json
// @Param id path string true "Test Case ID"
// @Success 200 {array} entities.TestCaseHistory
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-cases/{id}/history [get]
func (h *TestCaseHandler) GetTestCaseHistory(c *gin.Context) {
	testCaseID := c.Param("id")
	history, err := h.repo.GetTestCaseHistory(testCaseID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, history)
}
