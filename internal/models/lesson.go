package models

import "gorm.io/gorm"

type Lesson struct {
	gorm.Model

	SubjectId uint    `json:"subject_id"`
	Subject   Subject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title     string  `json:"title"`
}
