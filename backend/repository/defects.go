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

package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type defectRepository struct {
	db *gorm.DB
}

func NewDefectRepo(db *gorm.DB) DefectRepository {
	return &defectRepository{db: db}
}

func (r *defectRepository) GetDefectByID(id string) (entities.Defect, error) {
	var defect entities.Defect
	err := r.db.First(&defect, "id = ?", id).Error
	return defect, err
}

func (r *defectRepository) UpdateDefect(id string, req entities.CreateDefectRequest) (entities.Defect, error) {
	var defect entities.Defect
	if err := r.db.First(&defect, "id = ?", id).Error; err != nil {
		return defect, err
	}
	defect.ResultID = req.ResultID
	defect.ExternalLink = req.ExternalLink
	defect.Title = req.Title
	defect.Severity = req.Severity
	err := r.db.Save(&defect).Error
	return defect, err
}

func (r *defectRepository) DeleteDefect(id string) error {
	return r.db.Delete(&entities.Defect{}, "id = ?", id).Error
}
