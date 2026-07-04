package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
)

func orgnizationsRoutes(r *gin.RouterGroup) {
	r.GET("/organizations", endpoints.GetAllOrganizations)
	r.GET("/organizations/:id", endpoints.GetOrganizationByID)
	r.POST("/organizations", endpoints.CreateOrganization)
	r.DELETE("/organizations/:id", endpoints.DeleteOrganization)
}
