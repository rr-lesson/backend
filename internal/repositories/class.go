package repositories

import (
	"backend/internal/domains"
	"backend/internal/dto"
	"backend/internal/models"

	"gorm.io/gorm"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(
	db *gorm.DB,
) *ClassRepository {
	return &ClassRepository{
		db: db,
	}
}

func (r *ClassRepository) Create(data domains.Class) (*domains.Class, error) {
	class := data.ToModel()
	if err := r.db.Create(&class).Error; err != nil {
		return nil, err
	}

	return domains.FromClassModel(class), nil
}

func (r *ClassRepository) GetAll() (*[]dto.ClassDTO, error) {
	var classes []models.Class
	if err := r.db.Order("lower(name) asc").Find(&classes).Error; err != nil {
		return nil, err
	}

	result := make([]dto.ClassDTO, len(classes))
	for i, class := range classes {
		result[i] = dto.ClassDTO{
			Data: *domains.FromClassModel(&class),
		}
	}

	return &result, nil
}
