package handlers

import (
	"backend/internal/domains"
	"backend/internal/dto/requests"
	"backend/internal/dto/responses"
	"backend/internal/repositories"
	"backend/pkg/auth"
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authRepo *repositories.AuthRepository
}

func NewAuthHandler(
	authRepo *repositories.AuthRepository,
) *AuthHandler {
	return &AuthHandler{
		authRepo: authRepo,
	}
}

func (h *AuthHandler) RegisterRoutes(router fiber.Router) {
	g0 := router.Group("/auth")
	g0.Post("/login", h.login)
	g0.Post("/register", h.register)

	g1 := router.Group("/auth").Use(auth.AuthMiddleware())
	g1.Delete("/logout", h.logout)
}

// @id 					Login
// @tags 				auth
// @accept 			json
// @produce 		json
// @param 			body body requests.Login true "body"
// @success 		200 {object} responses.Login
// @router 			/api/v1/auth/login [post]
func (h *AuthHandler) login(c *fiber.Ctx) error {
	var req requests.Login
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	user, token, err := h.authRepo.Login(domains.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Error{
				Message: "Alamat email atau kata sandi tidak valid!",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    *token,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		Secure:   os.Getenv("GO_ENV") == "production",
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
	})

	return c.Status(fiber.StatusOK).JSON(responses.Login{
		User: *user,
	})
}

// @id 					Register
// @tags 				auth
// @accept 			json
// @produce 		json
// @param 			body body requests.Register true "body"
// @success 		200 {object} responses.Register
// @router 			/api/v1/auth/register [post]
func (h *AuthHandler) register(c *fiber.Ctx) error {
	var req requests.Register
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	user, token, err := h.authRepo.Register(domains.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     "user",
	})
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Error{
				Message: "Alamat email sudah terdaftar!",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    *token,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		Secure:   os.Getenv("GO_ENV") == "production",
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
	})

	return c.Status(fiber.StatusOK).JSON(responses.Register{
		User: *user,
	})
}

// @id 					Logout
// @tags 				auth
// @accept 			json
// @produce 		json
// @success 		200 {object} responses.Logout
// @router 			/api/v1/auth/logout [delete]
func (h *AuthHandler) logout(c *fiber.Ctx) error {
	session := auth.GetCurrentSession(c, h.authRepo)
	if session == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Error{
			Message: "Anda tidak memiliki akses untuk melakukan aksi ini!",
		})
	}

	if err := h.authRepo.Logout(session.Token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	c.ClearCookie("authToken")

	return c.Status(fiber.StatusOK).JSON(responses.Logout{
		Message: "ok",
	})
}
