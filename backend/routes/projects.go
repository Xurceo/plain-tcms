package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
)

func projectRoutes(r *gin.RouterGroup) {
	r.GET("/projects", endpoints.GetProjects)
	r.GET("/projects/:id", endpoints.GetProjectByID)
	r.DELETE("/projects/:id", endpoints.DeleteProject)
	r.GET("organizations/:id/projects", endpoints.GetProjectsByOrgID)
	r.POST("organizations/:id/projects", endpoints.CreateProject)
}
