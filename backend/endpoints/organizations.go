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
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/entities"
	"github.com/xurceo/plain-tcms/repository"
)

type OrganizationHandler struct {
	repo repository.OrganizationRepository
}

func NewOrganizationHandler(repo repository.OrganizationRepository) *OrganizationHandler {
	return &OrganizationHandler{repo: repo}
}

// GetAllOrganizations godoc
// @Summary Get all organizations
// @Tags Organizations
// @Produce json
// @Success 200 {array} entities.Organization
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations [get]
func (h *OrganizationHandler) GetAllOrganizations(c *gin.Context) {
	handleGetAll(c, h.repo.GetAllOrganizations)
}

// GetOrganizationByID godoc
// @Summary Get organization by ID
// @Tags Organizations
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} entities.Organization
// @Failure 404 {object} entities.ErrorResponse
// @Router /organizations/{id} [get]
func (h *OrganizationHandler) GetOrganizationByID(c *gin.Context) {
	handleGetByID(c, h.repo.GetOrganizationByID)
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
func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	handleCreate(c, h.repo.CreateOrganization)
}

// DeleteOrganization godoc
// @Summary Delete organization
// @Tags Organizations
// @Param id path string true "Organization ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations/{id} [delete]
func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	handleDelete(c, h.repo.DeleteOrganization)
}

// GetMembers godoc
// @Summary Get organization members
// @Tags Organizations
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {array} entities.OrganizationMember
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations/{id}/members [get]
func (h *OrganizationHandler) GetMembers(c *gin.Context) {
	id := c.Param("id")
	members, err := h.repo.GetMembers(id)
	if err != nil {
		slog.Error("failed to get members", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, members)
}

// AddMember godoc
// @Summary Add member to organization
// @Tags Organizations
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Param member body entities.CreateOrganizationMemberRequest true "Member"
// @Success 201 {object} entities.OrganizationMember
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations/{id}/members [post]
func (h *OrganizationHandler) AddMember(c *gin.Context) {
	orgID := c.Param("id")
	var req entities.CreateOrganizationMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	member, err := h.repo.AddMember(orgID, req)
	if err != nil {
		slog.Error("failed to add member", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, member)
}

// RemoveMember godoc
// @Summary Remove member from organization
// @Tags Organizations
// @Param id path string true "Organization ID"
// @Param user_id path string true "User ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations/{id}/members/{user_id} [delete]
func (h *OrganizationHandler) RemoveMember(c *gin.Context) {
	orgID := c.Param("id")
	userID := c.Param("user_id")
	if err := h.repo.RemoveMember(orgID, userID); err != nil {
		slog.Error("failed to remove member", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
