package domains

import (
	"backend/internal/models"
	"time"
)

type Lesson struct {
	Id        uint      `json:"id"`
	SubjectId uint      `json:"subject_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @name Lesson

func FromLessonModel(m *models.Lesson) *Lesson {
	return &Lesson{
		Id:        m.ID,
		SubjectId: m.SubjectId,
		Title:     m.Title,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *Lesson) ToModel() *models.Lesson {
	return &models.Lesson{
		SubjectId: m.SubjectId,
		Title:     m.Title,
	}
}
