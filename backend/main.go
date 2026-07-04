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
