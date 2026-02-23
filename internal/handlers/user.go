package handlers

import (
	"backend/internal/dto/responses"
	"backend/internal/repositories"
	"backend/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepo   *repositories.UserRepository
	authHelper *auth.AuthHelper
}

func NewUserHandler(
	userRepo *repositories.UserRepository,
	authHelper *auth.AuthHelper,
) *UserHandler {
	return &UserHandler{
		userRepo:   userRepo,
		authHelper: authHelper,
	}
}

func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	g0 := router.Group("/users").Use(auth.AuthMiddleware())
	g0.Get("/", h.getAllUsers)
}

// @id 					GetAllUsers
// @tags 				user
// @accept 			json
// @produce 		json
// @success 		200 {object} responses.GetAllUsers
// @failure 		500 {object} responses.Error
// @router 			/api/v1/users [get]
func (h *UserHandler) getAllUsers(c *fiber.Ctx) error {
	session, err := h.authHelper.GetCurrentSession(c)
	if err != nil {
		return err
	}

	if err := h.authHelper.ValidateAdmin(c, session); err != nil {
		return err
	}

	res, err := h.userRepo.GetAll(repositories.UserFilter{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(responses.GetAllUsers{
		Items: *res,
	})
}
