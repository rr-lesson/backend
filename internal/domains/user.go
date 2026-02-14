package domains

import (
	"backend/internal/models"
	"time"
)

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @name User

func FromUserModel(m *models.User) *User {
	return &User{
		Id:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		Role:      m.Role,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *User) ToModel() *models.User {
	return &models.User{
		Name:     m.Name,
		Email:    m.Email,
		Password: m.Password,
		Role:     m.Role,
	}
}
