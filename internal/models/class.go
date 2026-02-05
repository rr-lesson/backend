package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model

	Name string `json:"name"`
}
