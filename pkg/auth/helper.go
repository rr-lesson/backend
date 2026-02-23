package auth

import (
	"backend/internal/dto"
	"backend/internal/repositories"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AuthHelper struct {
	authRepo *repositories.AuthRepository
}

func NewAuthHelper(
	authRepo *repositories.AuthRepository,
) *AuthHelper {
	return &AuthHelper{
		authRepo: authRepo,
	}
}

func (h *AuthHelper) GetCurrentSession(c *fiber.Ctx) (*dto.UserSessionDTO, error) {
	token := c.Cookies("authToken")

	session, err := h.authRepo.GetSession(repositories.AuthFilter{
		Token:       token,
		IncludeUser: true,
	})
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (h *AuthHelper) ValidateAdmin(c *fiber.Ctx, session *dto.UserSessionDTO) error {
	if session.User.Role != "admin" {
		return errors.New("unauthorized")
	}

	return nil
}

func (h *AuthHelper) ValidateMember(c *fiber.Ctx, session *dto.UserSessionDTO) error {
	if session.User.Role != "member" {
		return errors.New("unauthorized")
	}

	return nil
}
