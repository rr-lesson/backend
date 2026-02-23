package repositories

import (
	"backend/internal/domains"
	"backend/internal/dto"
	"backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserFilter struct {
}

func NewUserRepository(
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll(filter UserFilter) (*[]dto.UserDTO, error) {
	var users []models.User

	query := r.db

	if err := query.Order("role asc").Order("lower(name) asc").Find(&users).Error; err != nil {
		return nil, err
	}

	result := make([]dto.UserDTO, len(users))
	for i, user := range users {
		result[i] = dto.UserDTO{
			Data: *domains.FromUserModel(&user),
		}
	}

	return &result, nil
}

func (r *UserRepository) Get(userId uint) (*dto.UserDTO, error) {
	var user models.User

	query := r.db

	if err := query.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &dto.UserDTO{Data: *domains.FromUserModel(&user)}, nil
}

func (r *UserRepository) Update(userId uint, data domains.User) (*domains.User, error) {
	user := data.ToModel()

	if err := r.db.Where("id = ?", userId).Updates(user).First(&user).Error; err != nil {
		return nil, err
	}

	return domains.FromUserModel(user), nil
}
