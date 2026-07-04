package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/repository"
	"gorm.io/gorm"
)

func projectRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewProjectHandler(repository.NewProjectRepo(db))
	tcH := endpoints.NewTestCaseHandler(repository.NewTestCaseRepo(db))
	tsH := endpoints.NewTestSuiteHandler(repository.NewTestSuiteRepo(db))
	tpH := endpoints.NewTestPlanHandler(repository.NewTestPlanRepo(db))
	trH := endpoints.NewTestRunHandler(repository.NewTestRunRepo(db))

	r.GET("/projects", h.GetProjects)
	r.GET("/projects/:id", h.GetProjectByID)
	r.PUT("/projects/:id", h.UpdateProject)
	r.DELETE("/projects/:id", h.DeleteProject)

	r.GET("organizations/:id/projects", h.GetProjectsByOrgID)
	r.POST("organizations/:id/projects", h.CreateProject)

	r.GET("/projects/:id/test-suites", tsH.GetTestSuitesByProject)
	r.POST("/projects/:id/test-suites", tsH.CreateTestSuite)
	r.GET("/projects/:id/test-cases", tcH.GetTestCases)
	r.POST("/projects/:id/test-cases", tcH.CreateTestCase)
	r.GET("/projects/:id/test-plans", tpH.GetTestPlansByProject)
	r.POST("/projects/:id/test-plans", tpH.CreateTestPlan)
	r.GET("/projects/:id/test-runs", trH.GetTestRunsByProject)
	r.POST("/projects/:id/test-runs", trH.CreateTestRun)
}
