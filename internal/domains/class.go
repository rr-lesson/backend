package domains

import (
	"backend/internal/models"
	"time"
)

type Class struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @name Class

func FromClassModel(m *models.Class) *Class {
	return &Class{
		Id:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *Class) ToModel() *models.Class {
	return &models.Class{
		Name: m.Name,
	}
}
