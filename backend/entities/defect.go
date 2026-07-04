package entities

import "time"

type Defect struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ResultID     *string   `gorm:"column:result_id;type:uuid" json:"result_id"`
	ExternalLink *string   `gorm:"column:external_link" json:"external_link"`
	Title        string    `gorm:"not null" json:"title"`
	Severity     *string   `json:"severity"`
	Status       string    `gorm:"default:open" json:"status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateDefectRequest struct {
	ResultID     *string `json:"result_id"`
	ExternalLink *string `json:"external_link"`
	Title        string  `json:"title"`
	Severity     *string `json:"severity"`
}

func (Defect) TableName() string { return "defects" }
