package auth

import (
	"backend/internal/domains"
	"backend/internal/repositories"

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

func (h *AuthHelper) GetCurrentSession(c *fiber.Ctx) *domains.UserSession {
	token := c.Cookies("authToken")

	session, err := h.authRepo.GetSession(repositories.AuthFilter{Token: token})
	if err != nil {
		return nil
	}

	return session
}
