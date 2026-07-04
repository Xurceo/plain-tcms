package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/entities"
	"github.com/xurceo/plain-tcms/repository"
)

type TestPlanHandler struct {
	repo repository.TestPlanRepository
}

func NewTestPlanHandler(repo repository.TestPlanRepository) *TestPlanHandler {
	return &TestPlanHandler{repo: repo}
}

// GetTestPlansByProject godoc
// @Summary Get test plans by project
// @Tags Test Plans
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {array} entities.TestPlan
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-plans [get]
func (h *TestPlanHandler) GetTestPlansByProject(c *gin.Context) {
	projectID := c.Param("id")
	plans, err := h.repo.GetTestPlansByProject(projectID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, plans)
}

// GetTestPlanByID godoc
// @Summary Get test plan by ID
// @Tags Test Plans
// @Produce json
// @Param id path string true "Test Plan ID"
// @Success 200 {object} entities.TestPlan
// @Failure 404 {object} entities.ErrorResponse
// @Router /test-plans/{id} [get]
func (h *TestPlanHandler) GetTestPlanByID(c *gin.Context) {
	handleGetByID(c, h.repo.GetTestPlanByID)
}

// CreateTestPlan godoc
// @Summary Create test plan
// @Tags Test Plans
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param test_plan body entities.CreateTestPlanRequest true "Test Plan"
// @Success 201 {object} entities.TestPlan
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id}/test-plans [post]
func (h *TestPlanHandler) CreateTestPlan(c *gin.Context) {
	projectID := c.Param("id")
	var req entities.CreateTestPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	plan, err := h.repo.CreateTestPlan(projectID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, plan)
}

// UpdateTestPlan godoc
// @Summary Update test plan
// @Tags Test Plans
// @Accept json
// @Produce json
// @Param id path string true "Test Plan ID"
// @Param test_plan body entities.CreateTestPlanRequest true "Test Plan"
// @Success 200 {object} entities.TestPlan
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-plans/{id} [put]
func (h *TestPlanHandler) UpdateTestPlan(c *gin.Context) {
	handleUpdate(c, h.repo.UpdateTestPlan)
}

// DeleteTestPlan godoc
// @Summary Delete test plan
// @Tags Test Plans
// @Param id path string true "Test Plan ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-plans/{id} [delete]
func (h *TestPlanHandler) DeleteTestPlan(c *gin.Context) {
	handleDelete(c, h.repo.DeleteTestPlan)
}
