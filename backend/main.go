package main

import (
	"log/slog"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xurceo/plain-tcms/db"
	_ "github.com/xurceo/plain-tcms/docs"
	"github.com/xurceo/plain-tcms/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	if err := db.Connect(); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	r := gin.Default()
	r.Use(cors.Default())

	routes.Setup(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// @title Plain-TCMS API
	// @version 1.0
	// @host localhost:8080
	// @BasePath /api/v1
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
