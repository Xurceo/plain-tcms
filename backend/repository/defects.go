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
