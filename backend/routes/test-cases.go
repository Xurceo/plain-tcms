package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/endpoints"
)

func testCasesRoutes(r *gin.RouterGroup) {
	r.GET("/projects/:id/test-cases", endpoints.GetTestCases)
	r.POST("/projects/:id/test-cases", endpoints.CreateTestCase)
}
