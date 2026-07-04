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

package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/repository"
)

// GetAllOrganizations godoc
// @Summary Get all organizations
// @Tags Organizations
// @Produce json
// @Success 200 {array} entities.Organization
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations [get]
func GetAllOrganizations(c *gin.Context) {
	handleGetAll(c, repository.GetAllOrganizations)
}

// GetOrganizationByID godoc
// @Summary Get organization by ID
// @Tags Organizations
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} entities.Organization
// @Failure 404 {object} entities.ErrorResponse
// @Router /organizations/{id} [get]
func GetOrganizationByID(c *gin.Context) {
	handleGetByID(c, repository.GetOrganizationByID)
}

// CreateOrganization godoc
// @Summary Create organization
// @Tags Organizations
// @Accept json
// @Produce json
// @Param organization body entities.CreateOrganizationRequest true "Organization"
// @Success 201 {object} entities.Organization
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations [post]
func CreateOrganization(c *gin.Context) {
	handleCreate(c, repository.CreateOrganization)
}

// DeleteOrganization godoc
// @Summary Delete organization
// @Tags Organizations
// @Param id path string true "Organization ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations/{id} [delete]
func DeleteOrganization(c *gin.Context) {
	handleDelete(c, repository.DeleteOrganization)
}
