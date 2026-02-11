package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model

	SubjectId uint    `json:"subject_id"`
	Subject   Subject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Question  string  `json:"question"`
}
