package repositories

import (
	"backend/internal/domains"
	"backend/internal/dto"
	"backend/internal/models"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

type QuestionFilter struct {
	IncludeUser    bool
	IncludeSubject bool
	IncludeClass   bool
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

func (r *QuestionRepository) GetAll(filter QuestionFilter) (*[]dto.QuestionDTO, error) {
	var questions []models.Question

	query := r.db

	if filter.IncludeUser {
		query = query.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		})
	}

	if filter.IncludeSubject {
		query = query.Preload("Subject")
	}

	if filter.IncludeClass {
		query = query.Preload("Subject.Class")
	}

	if err := query.Find(&questions).Error; err != nil {
		return nil, err
	}

	result := make([]dto.QuestionDTO, len(questions))
	for i, question := range questions {
		result[i] = dto.QuestionDTO{
			User:    *domains.FromUserModel(&question.User),
			Subject: *domains.FromSubjectModel(&question.Subject),
			Class:   *domains.FromClassModel(&question.Subject.Class),
			Data:    *domains.FromQuestionModel(&question),
		}
	}

	return &result, nil
}
