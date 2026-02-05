package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model

	ClassId uint   `json:"class_id"`
	Class   Class  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name    string `json:"name" gorm:"unique"`
}
