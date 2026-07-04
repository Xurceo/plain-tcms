package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/repository"
	"gorm.io/gorm"
)

func testRunRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewTestRunHandler(repository.NewTestRunRepo(db))

	r.GET("/test-runs/:id", h.GetTestRunByID)
	r.PUT("/test-runs/:id", h.UpdateTestRun)
	r.DELETE("/test-runs/:id", h.DeleteTestRun)

	r.GET("/test-runs/:id/cases", h.GetCasesByRun)
	r.POST("/test-runs/:id/cases", h.AddCaseToRun)
	r.DELETE("/test-runs/:id/cases/:test_case_id", h.RemoveCaseFromRun)

	r.GET("/test-runs/:id/results", h.GetResultsByRun)
	r.POST("/test-runs/:id/results", h.AddResultToRun)
}
