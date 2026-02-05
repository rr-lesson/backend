package repositories

import (
	"backend/internal/domains"
	"backend/internal/dto"
	"backend/internal/models"

	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{
		db: db,
	}
}

func (r *SubjectRepository) Create(data domains.Subject) (*domains.Subject, error) {
	subject := data.ToModel()
	if err := r.db.Create(&subject).Error; err != nil {
		return nil, err
	}

	return domains.FromSubjectModel(subject), nil
}

func (r *SubjectRepository) GetAll(classId uint) (*[]domains.Subject, error) {
	var subjects []models.Subject

	query := r.db
	if classId != 0 {
		query = query.Where("class_id = ?", classId)
	}

	if err := query.Order("lower(name) asc").Find(&subjects).Error; err != nil {
		return nil, err
	}

	result := make([]domains.Subject, len(subjects))
	for i, subject := range subjects {
		result[i] = *domains.FromSubjectModel(&subject)
	}

	return &result, nil
}

func (r *SubjectRepository) GetAllDetails() (*[]dto.SubjectDetail, error) {
	var subjects []models.Subject
	if err := r.db.Preload("Class").Order("lower(name) asc").Find(&subjects).Error; err != nil {
		return nil, err
	}

	result := make([]dto.SubjectDetail, len(subjects))
	for i, subject := range subjects {
		result[i] = dto.SubjectDetail{
			Subject: *domains.FromSubjectModel(&subject),
			Class:   *domains.FromClassModel(&subject.Class),
		}
	}

	return &result, nil
}
