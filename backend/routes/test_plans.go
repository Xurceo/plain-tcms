package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/repository"
	"gorm.io/gorm"
)

func testPlanRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewTestPlanHandler(repository.NewTestPlanRepo(db))

	r.GET("/test-plans/:id", h.GetTestPlanByID)
	r.PUT("/test-plans/:id", h.UpdateTestPlan)
	r.DELETE("/test-plans/:id", h.DeleteTestPlan)
}
