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

package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/xurceo/plain-tcms/repository"
)

type DefectHandler struct {
	repo repository.DefectRepository
}

func NewDefectHandler(repo repository.DefectRepository) *DefectHandler {
	return &DefectHandler{repo: repo}
}

// GetDefectByID godoc
// @Summary Get defect by ID
// @Tags Defects
// @Produce json
// @Param id path string true "Defect ID"
// @Success 200 {object} entities.Defect
// @Failure 404 {object} entities.ErrorResponse
// @Router /defects/{id} [get]
func (h *DefectHandler) GetDefectByID(c *gin.Context) {
	handleGetByID(c, h.repo.GetDefectByID)
}

// UpdateDefect godoc
// @Summary Update defect
// @Tags Defects
// @Accept json
// @Produce json
// @Param id path string true "Defect ID"
// @Param defect body entities.CreateDefectRequest true "Defect"
// @Success 200 {object} entities.Defect
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router /defects/{id} [put]
func (h *DefectHandler) UpdateDefect(c *gin.Context) {
	handleUpdate(c, h.repo.UpdateDefect)
}

// DeleteDefect godoc
// @Summary Delete defect
// @Tags Defects
// @Param id path string true "Defect ID"
// @Success 204
// @Failure 500 {object} entities.ErrorResponse
// @Router /defects/{id} [delete]
func (h *DefectHandler) DeleteDefect(c *gin.Context) {
	handleDelete(c, h.repo.DeleteDefect)
}
