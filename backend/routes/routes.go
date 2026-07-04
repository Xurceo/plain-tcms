package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")
	organizationRoutes(api, db)
	projectRoutes(api, db)
	testCasesRoutes(api, db)
	testSuiteRoutes(api, db)
	testPlanRoutes(api, db)
	testRunRoutes(api, db)
	testResultRoutes(api, db)
	defectRoutes(api, db)
	authRoutes(api, db)
}
