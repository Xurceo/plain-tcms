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

func organizationRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewOrganizationHandler(repository.NewOrganizationRepo(db))

	r.GET("/organizations", h.GetAllOrganizations)
	r.POST("/organizations", h.CreateOrganization)
	r.GET("/organizations/:id", h.GetOrganizationByID)
	r.DELETE("/organizations/:id", h.DeleteOrganization)
	r.GET("/organizations/:id/members", h.GetMembers)
	r.POST("/organizations/:id/members", h.AddMember)
	r.DELETE("/organizations/:id/members/:user_id", h.RemoveMember)
}
