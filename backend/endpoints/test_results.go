package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/entities"
	"github.com/xurceo/plain-tcms/repository"
)

type TestResultHandler struct {
	repo repository.TestResultRepository
}

func NewTestResultHandler(repo repository.TestResultRepository) *TestResultHandler {
	return &TestResultHandler{repo: repo}
}

// GetTestResultByID godoc
// @Summary Get test result by ID
// @Tags Test Results
// @Produce json
// @Param id path string true "Test Result ID"
// @Success 200 {object} entities.TestResult
// @Failure 404 {object} entities.ErrorResponse
// @Router /test-results/{id} [get]
func (h *TestResultHandler) GetTestResultByID(c *gin.Context) {
	handleGetByID(c, h.repo.GetTestResultByID)
}

// UpdateTestResult godoc
// @Summary Update test result
// @Tags Test Results
// @Accept json
// @Produce json
// @Param id path string true "Test Result ID"
// @Param result body entities.CreateTestResultRequest true "Test Result"
// @Success 200 {object} entities.TestResult
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-results/{id} [put]
func (h *TestResultHandler) UpdateTestResult(c *gin.Context) {
	handleUpdate(c, h.repo.UpdateTestResult)
}

// GetAttachmentsByResult godoc
// @Summary Get attachments for a test result
// @Tags Test Results
// @Produce json
// @Param id path string true "Test Result ID"
// @Success 200 {array} entities.ResultAttachment
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-results/{id}/attachments [get]
func (h *TestResultHandler) GetAttachmentsByResult(c *gin.Context) {
	resultID := c.Param("id")
	attachments, err := h.repo.GetAttachmentsByResult(resultID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, attachments)
}

// AddAttachmentToResult godoc
// @Summary Add attachment to test result
// @Tags Test Results
// @Accept json
// @Produce json
// @Param id path string true "Test Result ID"
// @Param attachment body entities.CreateResultAttachmentRequest true "Attachment"
// @Success 201 {object} entities.ResultAttachment
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /test-results/{id}/attachments [post]
func (h *TestResultHandler) AddAttachmentToResult(c *gin.Context) {
	resultID := c.Param("id")
	var req entities.CreateResultAttachmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	attachment, err := h.repo.AddAttachmentToResult(resultID, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, attachment)
}
