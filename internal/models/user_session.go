package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSession struct {
	gorm.Model

	UserId     uint      `json:"user_id"`
	User       User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Token      string    `json:"token" gorm:"unique"`
	LastUsedAt time.Time `json:"last_used_at"`
}
