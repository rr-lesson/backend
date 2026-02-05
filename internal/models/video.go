package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model

	LessonId    uint   `json:"lesson_id"`
	Lesson      Lesson `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FilePath    string `json:"file_path"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
