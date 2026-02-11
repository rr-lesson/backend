package repositories

import (
	"backend/internal/domains"
	"backend/internal/models"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

type QuestionFilter struct {
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

func (r *QuestionRepository) Create(data domains.Question) (*domains.Question, error) {
	question := data.ToModel()
	if err := r.db.Create(&question).Error; err != nil {
		return nil, err
	}

	return domains.FromQuestionModel(question), nil
}

func (r *QuestionRepository) GetAll() (*[]domains.Question, error) {
	var questions []models.Question

	query := r.db

	if err := query.Find(&questions).Error; err != nil {
		return nil, err
	}

	result := make([]domains.Question, len(questions))
	for i, question := range questions {
		result[i] = *domains.FromQuestionModel(&question)
	}

	return &result, nil
}
