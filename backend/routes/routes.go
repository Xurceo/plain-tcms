package routes

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	api := r.Group("/api/v1")
	projectRoutes(api)
	testCasesRoutes(api)
	orgnizationsRoutes(api)
}
