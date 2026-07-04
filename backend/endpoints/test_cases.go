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
