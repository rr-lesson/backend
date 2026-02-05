package domains

import (
	"backend/internal/models"
	"time"
)

type Video struct {
	Id          uint      `json:"id"`
	LessonId    uint      `json:"lesson_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FilePath    string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
} // @name Video

func FromVideoModel(m *models.Video) *Video {
	return &Video{
		Id:          m.ID,
		LessonId:    m.LessonId,
		Title:       m.Title,
		Description: m.Description,
		FilePath:    m.FilePath,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func (m *Video) ToModel() *models.Video {
	return &models.Video{
		LessonId:    m.LessonId,
		FilePath:    m.FilePath,
		Title:       m.Title,
		Description: m.Description,
	}
}
