package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model

	QuestionId uint     `json:"question_id"`
	Question   Question `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Answer     string   `json:"answer"`
}
