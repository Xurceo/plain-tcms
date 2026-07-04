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

func testCasesRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewTestCaseHandler(repository.NewTestCaseRepo(db))

	r.GET("/test-cases/:id", h.GetTestCaseByID)
	r.PUT("/test-cases/:id", h.UpdateTestCase)
	r.DELETE("/test-cases/:id", h.DeleteTestCase)
	r.GET("/test-cases/:id/history", h.GetTestCaseHistory)
}
