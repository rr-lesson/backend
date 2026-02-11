package domains

import (
	"backend/internal/models"
	"time"
)

type Question struct {
	Id        uint      `json:"id"`
	SubjectId uint      `json:"subject_id"`
	Question  string    `json:"question"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @name Question

func FromQuestionModel(m *models.Question) *Question {
	return &Question{
		Id:        m.ID,
		SubjectId: m.SubjectId,
		Question:  m.Question,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *Question) ToModel() *models.Question {
	return &models.Question{
		SubjectId: m.SubjectId,
		Question:  m.Question,
	}
}
