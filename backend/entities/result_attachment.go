package entities

import "time"

type ResultAttachment struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ResultID  string    `gorm:"column:result_id;type:uuid;not null" json:"result_id"`
	FileURL   string    `gorm:"column:file_url;not null" json:"file_url"`
	FileType  *string   `gorm:"column:file_type" json:"file_type"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateResultAttachmentRequest struct {
	FileURL  string  `json:"file_url"`
	FileType *string `json:"file_type"`
}

func (ResultAttachment) TableName() string { return "result_attachments" }
