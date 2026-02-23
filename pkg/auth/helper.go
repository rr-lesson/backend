package auth

import (
	"backend/internal/dto"
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

func (h *AuthHelper) GetCurrentSession(c *fiber.Ctx) (*dto.UserSessionDTO, error) {
	token := c.Cookies("authToken")

	session, err := h.authRepo.GetSession(repositories.AuthFilter{
		Token:       token,
		IncludeUser: true})
	if err != nil {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Error{
			Message: "Anda tidak terautentikasi!",
		})
	}

	return session, nil
}

func (h *AuthHelper) ValidateAdmin(c *fiber.Ctx, session *dto.UserSessionDTO) error {
	if session.User.Role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Error{
			Message: "Anda tidak memiliki akses untuk melakukan ini!",
		})
	}

	return nil
}

func (h *AuthHelper) ValidateMember(c *fiber.Ctx, session *dto.UserSessionDTO) error {
	if session.User.Role != "member" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Error{
			Message: "Anda tidak memiliki akses untuk melakukan ini!",
		})
	}

	return nil
}
