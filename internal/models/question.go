package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model

	UserId    uint    `json:"user_id"`
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubjectId uint    `json:"subject_id"`
	Subject   Subject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Question  string  `json:"question"`

	// back references
	Attachments []QuestionAttachment
}
