package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/repository"
	"gorm.io/gorm"
)

func authRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := endpoints.NewAuthHandler(repository.NewUserRepo(db))

	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)
}
