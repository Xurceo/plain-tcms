package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/repository"
	"gorm.io/gorm"
)

func testSuiteRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewTestSuiteHandler(repository.NewTestSuiteRepo(db))

	r.GET("/test-suites/:id", h.GetTestSuiteByID)
	r.PUT("/test-suites/:id", h.UpdateTestSuite)
	r.DELETE("/test-suites/:id", h.DeleteTestSuite)
}
