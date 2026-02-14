package models

import "gorm.io/gorm"

type QuestionAttachment struct {
	gorm.Model

	QuestionId uint     `json:"question_id"`
	Question   Question `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Path       string   `json:"path"`
}
