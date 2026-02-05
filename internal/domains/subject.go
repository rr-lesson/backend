package domains

import (
	"backend/internal/models"
	"time"
)

type Subject struct {
	Id        uint      `json:"id"`
	ClassId   uint      `json:"class_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @name Subject

func FromSubjectModel(m *models.Subject) *Subject {
	return &Subject{
		Id:        m.ID,
		ClassId:   m.ClassId,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *Subject) ToModel() *models.Subject {
	return &models.Subject{
		ClassId: m.ClassId,
		Name:    m.Name,
	}
}
