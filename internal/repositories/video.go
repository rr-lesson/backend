package repositories

import (
	"backend/internal/domains"
	"backend/internal/dto"
	"backend/internal/models"

	"gorm.io/gorm"
)

type VideoFilter struct {
	LessonId uint
}

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(
	db *gorm.DB,
) *VideoRepository {
	return &VideoRepository{
		db: db,
	}
}

func (r *VideoRepository) Create(data domains.Video) (*domains.Video, error) {
	video := data.ToModel()
	if err := r.db.Create(&video).Error; err != nil {
		return nil, err
	}

	return domains.FromVideoModel(video), nil
}

func (r *VideoRepository) GetAll(filter VideoFilter) (*[]domains.Video, error) {
	var videos []models.Video

	query := r.db
	if filter.LessonId != 0 {
		query = query.Where("lesson_id = ?", filter.LessonId)
	}

	if err := query.Order("lower(title) asc").Find(&videos).Error; err != nil {
		return nil, err
	}

	result := make([]domains.Video, len(videos))
	for i, video := range videos {
		result[i] = *domains.FromVideoModel(&video)
	}

	return &result, nil
}

func (r *VideoRepository) GetAllByLessonId(lessonId uint) (*[]domains.Video, error) {
	var videos []models.Video
	if err := r.db.Where("lesson_id = ?", lessonId).Order("lower(title) asc").Find(&videos).Error; err != nil {
		return nil, err
	}

	var result []domains.Video
	for _, video := range videos {
		result = append(result, *domains.FromVideoModel(&video))
	}

	return &result, nil
}

func (r *VideoRepository) GetAllWithDetail() (*[]dto.VideoDetail, error) {
	var videos []models.Video
	if err := r.db.Preload("Lesson.Subject.Class").Order("lower(title) asc").Find(&videos).Error; err != nil {
		return nil, err
	}

	result := make([]dto.VideoDetail, len(videos))
	for i, video := range videos {
		result[i] = dto.VideoDetail{
			Video:   *domains.FromVideoModel(&video),
			Lesson:  *domains.FromLessonModel(&video.Lesson),
			Subject: *domains.FromSubjectModel(&video.Lesson.Subject),
			Class:   *domains.FromClassModel(&video.Lesson.Subject.Class),
		}
	}

	return &result, nil
}

func (r *VideoRepository) GetWithDetail(videoId uint) (*dto.VideoDetail, error) {
	var video models.Video
	if err := r.db.Preload("Lesson.Subject.Class").First(&video, videoId).Error; err != nil {
		return nil, err
	}

	return &dto.VideoDetail{
		Video:   *domains.FromVideoModel(&video),
		Lesson:  *domains.FromLessonModel(&video.Lesson),
		Subject: *domains.FromSubjectModel(&video.Lesson.Subject),
		Class:   *domains.FromClassModel(&video.Lesson.Subject.Class),
	}, nil
}
