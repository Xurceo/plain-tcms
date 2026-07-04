package endpoints

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/entities"
	"github.com/xurceo/plain-tcms/repository"
)

// GetProjects godoc
// @Summary Get all projects
// @Tags Projects
// @Produce json
// @Success 200 {array} entities.Project
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects [get]
func GetProjects(c *gin.Context) {
	projects, err := repository.GetAllProjects()
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
func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	project, err := repository.GetProjectByID(id)
	if err != nil {
		slog.Error("failed to get project by id", "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "project not found"})
		return
	}
	c.JSON(http.StatusOK, project)
}

// GetProjectsByOrgID godoc
// @Summary Get projects by organization
// @Tags Projects
// @Produce json
// @Param org_id path string true "Organization ID"
// @Success 200 {array} entities.Project
// @Failure 500 {object} entities.ErrorResponse
// @Router /organizations/{org_id}/projects [get]
func GetProjectsByOrgID(c *gin.Context) {
	orgID := c.Param("id")
	projects, err := repository.GetProjectsByOrgID(orgID)
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
func CreateProject(c *gin.Context) {
	orgID := c.Param("id")

	var req entities.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.OrgID = orgID

	project, err := repository.CreateProject(req)
	if err != nil {
		slog.Error("failed to create project", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, project)
}

// DeleteProject godoc
// @Summary Delete project
// @Tags Projects
// @Param id path string true "Project ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /projects/{id} [delete]
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DeleteProject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
