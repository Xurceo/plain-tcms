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
