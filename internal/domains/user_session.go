package domains

import (
	"backend/internal/models"
	"time"
)

type UserSession struct {
	Id         uint      `json:"id"`
	UserId     uint      `json:"user_id"`
	Token      string    `json:"token"`
	LastUsedAt time.Time `json:"last_used_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
} // @name UserSession

func FromUserSessionModel(m *models.UserSession) *UserSession {
	return &UserSession{
		Id:         m.ID,
		UserId:     m.UserId,
		Token:      m.Token,
		LastUsedAt: m.LastUsedAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func (m *UserSession) ToModel() *models.UserSession {
	return &models.UserSession{
		UserId:     m.UserId,
		Token:      m.Token,
		LastUsedAt: m.LastUsedAt,
	}
}
