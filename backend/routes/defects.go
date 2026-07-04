package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/repository"
	"gorm.io/gorm"
)

func defectRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewDefectHandler(repository.NewDefectRepo(db))

	r.GET("/defects/:id", h.GetDefectByID)
	r.PUT("/defects/:id", h.UpdateDefect)
	r.DELETE("/defects/:id", h.DeleteDefect)
}
