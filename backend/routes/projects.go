/*
 * // TCMS - Test Case Management System
 * // Copyright (C) 2026 Pavlo Shnal
 * //
 * // This program is free software: you can redistribute it and/or modify
 * // it under the terms of the GNU Affero General Public License as published
 * // by the Free Software Foundation, either version 3 of the License, or
 * // (at your option) any later version.
 * //
 * // This program is distributed in the hope that it will be useful,
 * // but WITHOUT ANY WARRANTY; without even the implied warranty of
 * // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * // GNU Affero General Public License for more details.
 * //
 * // You should have received a copy of the GNU Affero General Public License
 * // along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

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
