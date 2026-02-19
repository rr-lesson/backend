package repositories

import (
	"backend/internal/domains"
	"backend/internal/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

type AuthFilter struct {
	Token string
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Login(data domains.User) (*domains.User, *string, error) {
	var user models.User
	if err := r.db.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return nil, nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, nil, err
	}

	token := uuid.NewString()
	if err := r.db.Create(&models.UserSession{
		UserId:     user.ID,
		Token:      token,
		LastUsedAt: time.Now(),
	}).Error; err != nil {
		return nil, nil, err
	}

	user.Password = ""

	return domains.FromUserModel(&user), &token, nil
}

func (r *AuthRepository) Register(data domains.User) (*domains.User, *string, error) {
	user := data.ToModel()
	token := uuid.NewString()

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)

		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if err := tx.Create(&models.UserSession{
			UserId:     user.ID,
			Token:      token,
			LastUsedAt: time.Now(),
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, nil, err
	}

	user.Password = ""

	return domains.FromUserModel(user), &token, nil
}

func (r *AuthRepository) RefreshToken(oldToken, newToken string) (*domains.User, error) {
	var userSession models.UserSession
	if err := r.db.Where("token = ?", oldToken).Preload("User").First(&userSession).Error; err != nil {
		return nil, err
	}

	userSession.Token = newToken
	userSession.LastUsedAt = time.Now()

	if err := r.db.Save(&userSession).Error; err != nil {
		return nil, err
	}

	return domains.FromUserModel(&userSession.User), nil
}

func (r *AuthRepository) Logout(token string) error {
	if err := r.db.Where("token = ?", token).Unscoped().Delete(&models.UserSession{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetSession(filter AuthFilter) (*domains.UserSession, error) {
	var userSession models.UserSession
	if err := r.db.Where("token = ?", filter.Token).First(&userSession).Error; err != nil {
		return nil, err
	}

	return domains.FromUserSessionModel(&userSession), nil
}
