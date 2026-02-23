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
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authRepo   *repositories.AuthRepository
	authHelper *auth.AuthHelper
}

func NewAuthHandler(
	authRepo *repositories.AuthRepository,
	authHelper *auth.AuthHelper,
) *AuthHandler {
	return &AuthHandler{
		authRepo:   authRepo,
		authHelper: authHelper,
	}
}

func (h *AuthHandler) RegisterRoutes(router fiber.Router) {
	g0 := router.Group("/auth")
	g0.Post("/login", h.login)
	g0.Post("/register", h.register)

	g1 := router.Group("/auth").Use(auth.AuthMiddleware())
	g1.Put("/refresh", h.refreshToken)
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

	cookie := fiber.Cookie{
		Name:     "authToken",
		Value:    *token,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		Secure:   false,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
	}

	if os.Getenv("GO_ENV") == "production" {
		cookie.Secure = true
		cookie.Domain = ".rizalanggoro.my.id"
	}

	c.Cookie(&cookie)

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
		Role:     "guest",
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

	cookie := fiber.Cookie{
		Name:     "authToken",
		Value:    *token,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		Secure:   false,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
	}

	if os.Getenv("GO_ENV") == "production" {
		cookie.Secure = true
		cookie.Domain = ".rizalanggoro.my.id"
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(responses.Register{
		User: *user,
	})
}

// @id 					RefreshToken
// @tags 				auth
// @accept 			json
// @produce 		json
// @success 		200 {object} responses.RefreshToken
// @failure 		500 {object} responses.Error
// @router 			/api/v1/auth/refresh [put]
func (h *AuthHandler) refreshToken(c *fiber.Ctx) error {
	session, err := h.authHelper.GetCurrentSession(c)
	if err != nil {
		return err
	}

	newToken := uuid.NewString()
	user, err := h.authRepo.RefreshToken(session.Data.Token, newToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.Error{
				Code:    fiber.StatusUnauthorized,
				Message: err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "authToken",
		Value:    newToken,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		Secure:   false,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteLaxMode,
	}

	if os.Getenv("GO_ENV") == "production" {
		cookie.Secure = true
		cookie.Domain = ".rizalanggoro.my.id"
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(responses.RefreshToken{
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
	session, err := h.authHelper.GetCurrentSession(c)
	if err != nil {
		return err
	}

	if err := h.authRepo.Logout(session.Data.Token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	c.ClearCookie("authToken")

	return c.Status(fiber.StatusOK).JSON(responses.Logout{
		Message: "ok",
	})
}
