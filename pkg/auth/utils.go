package auth

import (
	"backend/internal/domains"
	"backend/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentSession(c *fiber.Ctx, authRepo *repositories.AuthRepository) *domains.UserSession {
	token := c.Cookies("authToken")

	session, err := authRepo.GetSession(repositories.AuthFilter{Token: token})
	if err != nil {
		return nil
	}

	return session
}
