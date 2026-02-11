package domains

import (
	"backend/internal/models"
	"time"
)

type Answer struct {
	Id         uint      `json:"id"`
	QuestionId uint      `json:"question_id"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
} // @name Answer

func FromAnswerModel(m *models.Answer) *Answer {
	return &Answer{
		Id:         m.ID,
		QuestionId: m.QuestionId,
		Answer:     m.Answer,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func (m *Answer) ToModel() *models.Answer {
	return &models.Answer{
		QuestionId: m.QuestionId,
		Answer:     m.Answer,
	}
}
