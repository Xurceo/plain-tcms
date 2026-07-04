package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/repository"
	"gorm.io/gorm"
)

func testResultRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewTestResultHandler(repository.NewTestResultRepo(db))

	r.GET("/test-results/:id", h.GetTestResultByID)
	r.PUT("/test-results/:id", h.UpdateTestResult)

	r.GET("/test-results/:id/attachments", h.GetAttachmentsByResult)
	r.POST("/test-results/:id/attachments", h.AddAttachmentToResult)
}
