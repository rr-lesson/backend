package repositories

import (
	"backend/internal/domains"
	"backend/internal/dto"
	"backend/internal/models"
	"context"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db    *gorm.DB
	minio *minio.Client
}

type QuestionFilter struct {
	QuestionId         uint
	Keyword            string
	IncludeUser        bool
	IncludeSubject     bool
	IncludeClass       bool
	IncludeAttachments bool
}

func NewQuestionRepository(db *gorm.DB, minio *minio.Client) *QuestionRepository {
	return &QuestionRepository{
		db:    db,
		minio: minio,
	}
}

func (r *QuestionRepository) Create(ctx context.Context, data domains.Question, images []*multipart.FileHeader) (*domains.Question, error) {
	question := data.ToModel()
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&question).Error; err != nil {
			return err
		}

		if len(images) > 0 {
			for _, image := range images {
				src, err := image.Open()
				if err != nil {
					return err
				}
				defer src.Close()

				objectName := path.Join("public", "question-attachments", uuid.NewString()+filepath.Ext(image.Filename))
				if _, err := r.minio.PutObject(
					ctx,
					os.Getenv("MINIO_BUCKET"),
					objectName,
					src,
					image.Size,
					minio.PutObjectOptions{},
				); err != nil {
					return err
				}

				if err := tx.Create(&models.QuestionAttachment{
					QuestionId: question.ID,
					Path:       objectName,
				}).Error; err != nil {
					return err
				}
			}
		}

		return nil
	}); err != nil {
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

	if filter.Keyword != "" {
		query = query.Where("lower(question) LIKE lower(?)", "%"+filter.Keyword+"%")
	}

	if filter.IncludeAttachments {
		query = query.Preload("Attachments")
	}

	if err := query.Find(&questions).Error; err != nil {
		return nil, err
	}

	result := make([]dto.QuestionDTO, len(questions))
	for i, question := range questions {
		attachments := make([]domains.QuestionAttachment, len(question.Attachments))
		for j, attachment := range question.Attachments {
			attachments[j] = *domains.FromQuestionAttachmentModel(&attachment)
		}

		result[i] = dto.QuestionDTO{
			User:        *domains.FromUserModel(&question.User),
			Subject:     *domains.FromSubjectModel(&question.Subject),
			Class:       *domains.FromClassModel(&question.Subject.Class),
			Data:        *domains.FromQuestionModel(&question),
			Attachments: attachments,
		}
	}

	return &result, nil
}

func (r *QuestionRepository) Get(filter QuestionFilter) (*dto.QuestionDTO, error) {
	var question models.Question

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

	if filter.IncludeAttachments {
		query = query.Preload("Attachments")
	}

	if err := query.First(&question, filter.QuestionId).Error; err != nil {
		return nil, err
	}

	attachments := make([]domains.QuestionAttachment, len(question.Attachments))
	for i, attachment := range question.Attachments {
		attachments[i] = *domains.FromQuestionAttachmentModel(&attachment)

		url := url.URL{
			Scheme: os.Getenv("MINIO_SCHEME"),
			Host:   os.Getenv("MINIO_ENDPOINT"),
			Path: path.Join(
				os.Getenv("MINIO_BUCKET"),
				attachments[i].Path,
			),
		}
		attachments[i].Path = url.String()
	}

	result := dto.QuestionDTO{
		User:        *domains.FromUserModel(&question.User),
		Subject:     *domains.FromSubjectModel(&question.Subject),
		Class:       *domains.FromClassModel(&question.Subject.Class),
		Data:        *domains.FromQuestionModel(&question),
		Attachments: attachments,
	}

	return &result, nil
}
