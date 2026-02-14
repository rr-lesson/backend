package domains

import (
	"backend/internal/models"
	"time"
)

type QuestionAttachment struct {
	Id         uint      `json:"id"`
	QuestionId uint      `json:"question_id"`
	Path       string    `json:"path"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
} // @name QuestionAttachment

func FromQuestionAttachmentModel(m *models.QuestionAttachment) *QuestionAttachment {
	return &QuestionAttachment{
		Id:         m.ID,
		QuestionId: m.QuestionId,
		Path:       m.Path,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func (m *QuestionAttachment) ToModel() *models.QuestionAttachment {
	return &models.QuestionAttachment{
		QuestionId: m.Id,
		Path:       m.Path,
	}
}
