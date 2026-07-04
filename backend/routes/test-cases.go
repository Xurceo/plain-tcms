package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/repository"
	"gorm.io/gorm"
)

func testCasesRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewTestCaseHandler(repository.NewTestCaseRepo(db))

	r.GET("/test-cases/:id", h.GetTestCaseByID)
	r.PUT("/test-cases/:id", h.UpdateTestCase)
	r.DELETE("/test-cases/:id", h.DeleteTestCase)
	r.GET("/test-cases/:id/history", h.GetTestCaseHistory)
}
