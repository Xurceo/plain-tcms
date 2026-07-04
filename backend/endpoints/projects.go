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

type ProjectHandler struct {
	repo repository.ProjectRepository
}

func NewProjectHandler(repo repository.ProjectRepository) *ProjectHandler {
	return &ProjectHandler{repo: repo}
}

// GetProjects godoc
// @Summary Get all projects
// @Tags Projects
// @Produce json
// @Success 200 {array} entities.Project
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects [get]
func (h *ProjectHandler) GetProjects(c *gin.Context) {
	projects, err := h.repo.GetAllProjects()
	if err != nil {
		slog.Error("failed to get projects", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// GetProjectByID godoc
// @Summary Get project by ID
// @Tags Projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} entities.Project
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id} [get]
func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	handleGetByID(c, h.repo.GetProjectByID)
}

// GetProjectsByOrgID godoc
// @Summary Get projects by organization
// @Tags Projects
// @Produce json
// @Param org_id path string true "Organization ID"
// @Success 200 {array} entities.Project
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations/{org_id}/projects [get]
func (h *ProjectHandler) GetProjectsByOrgID(c *gin.Context) {
	orgID := c.Param("id")
	projects, err := h.repo.GetProjectsByOrgID(orgID)
	if err != nil {
		slog.Error("failed to get projects by org", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// CreateProject godoc
// @Summary Create project
// @Tags Projects
// @Accept json
// @Produce json
// @Param org_id path string true "Organization ID"
// @Param project body entities.CreateProjectRequest true "Project"
// @Success 201 {object} entities.Project
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations/{org_id}/projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	orgID := c.Param("id")

	var req entities.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.OrgID = orgID

	project, err := h.repo.CreateProject(req)
	if err != nil {
		slog.Error("failed to create project", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, project)
}

// UpdateProject godoc
// @Summary Update project
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body entities.CreateProjectRequest true "Project"
// @Success 200 {object} entities.Project
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id} [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	handleUpdate(c, h.repo.UpdateProject)
}

// DeleteProject godoc
// @Summary Delete project
// @Tags Projects
// @Param id path string true "Project ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteProject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
