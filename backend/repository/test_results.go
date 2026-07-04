package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type testResultRepository struct {
	db *gorm.DB
}

func NewTestResultRepo(db *gorm.DB) TestResultRepository {
	return &testResultRepository{db: db}
}

func (r *testResultRepository) GetTestResultByID(id string) (entities.TestResult, error) {
	var result entities.TestResult
	err := r.db.First(&result, "id = ?", id).Error
	return result, err
}

func (r *testResultRepository) UpdateTestResult(id string, req entities.CreateTestResultRequest) (entities.TestResult, error) {
	var result entities.TestResult
	if err := r.db.First(&result, "id = ?", id).Error; err != nil {
		return result, err
	}
	result.TestCaseID = req.TestCaseID
	result.Status = req.Status
	result.Comment = req.Comment
	result.DurationMs = req.DurationMs
	err := r.db.Save(&result).Error
	return result, err
}

func (r *testResultRepository) GetAttachmentsByResult(resultID string) ([]entities.ResultAttachment, error) {
	var attachments []entities.ResultAttachment
	err := r.db.Where("result_id = ?", resultID).Find(&attachments).Error
	return attachments, err
}

func (r *testResultRepository) AddAttachmentToResult(resultID string, req entities.CreateResultAttachmentRequest) (entities.ResultAttachment, error) {
	attachment := entities.ResultAttachment{
		ResultID: resultID,
		FileURL:  req.FileURL,
		FileType: req.FileType,
	}
	err := r.db.Create(&attachment).Error
	return attachment, err
}
