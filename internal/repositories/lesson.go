package repositories

import (
	"backend/internal/domains"
	"backend/internal/dto"
	"backend/internal/models"

	"gorm.io/gorm"
)

type LessonFilter struct {
	ClassId   uint
	SubjectId uint
}

type LessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(
	db *gorm.DB,
) *LessonRepository {
	return &LessonRepository{
		db: db,
	}
}

func (r *LessonRepository) Create(data domains.Lesson) (*domains.Lesson, error) {
	lesson := data.ToModel()
	if err := r.db.Create(&lesson).Error; err != nil {
		return nil, err
	}

	return domains.FromLessonModel(lesson), nil
}

func (r *LessonRepository) GetAll(filter LessonFilter) (*[]domains.Lesson, error) {
	var lessons []models.Lesson

	query := r.db
	if filter.ClassId != 0 {
		query = query.Joins("join subjects on subjects.id = lessons.subject_id").Where("subjects.class_id = ?", filter.ClassId)
	}
	if filter.SubjectId != 0 {
		query = query.Where("lessons.subject_id = ?", filter.SubjectId)
	}

	if err := query.Order("lower(title) asc").Find(&lessons).Error; err != nil {
		return nil, err
	}

	result := make([]domains.Lesson, len(lessons))
	for i, lesson := range lessons {
		result[i] = *domains.FromLessonModel(&lesson)
	}

	return &result, nil
}

func (r *LessonRepository) GetAllBySubjectId(subjectId uint) (*[]domains.Lesson, error) {
	var lessons []models.Lesson
	if err := r.db.Where("subject_id = ?", subjectId).Order("lower(title) asc").Find(&lessons).Error; err != nil {
		return nil, err
	}

	result := make([]domains.Lesson, len(lessons))
	for i, lesson := range lessons {
		result[i] = *domains.FromLessonModel(&lesson)
	}

	return &result, nil
}

func (r *LessonRepository) GetAllWithClassSubject() (*[]dto.LessonClassSubject, error) {
	var lessons []models.Lesson
	if err := r.db.Preload("Subject.Class").Find(&lessons).Error; err != nil {
		return nil, err
	}

	result := make([]dto.LessonClassSubject, len(lessons))
	for i, lesson := range lessons {
		result[i] = dto.LessonClassSubject{
			Lesson:  *domains.FromLessonModel(&lesson),
			Class:   *domains.FromClassModel(&lesson.Subject.Class),
			Subject: *domains.FromSubjectModel(&lesson.Subject),
		}
	}

	return &result, nil
}
