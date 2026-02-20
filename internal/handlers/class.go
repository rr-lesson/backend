package handlers

import (
	"backend/internal/domains"
	"backend/internal/dto/requests"
	"backend/internal/dto/responses"
	"backend/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

type ClassHandler struct {
	classRepo *repositories.ClassRepository
}

func NewClassHandler(
	classRepo *repositories.ClassRepository,
) *ClassHandler {
	return &ClassHandler{
		classRepo: classRepo,
	}
}

func (h *ClassHandler) RegisterRoutes(router fiber.Router) {
	g0 := router.Group("/classes")
	g0.Post("/", h.createClass)
	g0.Get("/", h.getAllClasses)
}

// @id 					CreateClass
// @tags 				class
// @accept 			json
// @produce 		json
// @param 			body body requests.CreateClass true "body"
// @success 		200 {object} responses.CreateClass
// @router 			/api/v1/classes [post]
func (h *ClassHandler) createClass(c *fiber.Ctx) error {
	var req requests.CreateClass
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.classRepo.Create(domains.Class{
		Name: req.Name,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.CreateClass{
		Class: *res,
	})
}

// @id 					GetAllClasses
// @tags 				class
// @accept 			json
// @produce 		json
// @success 		200 {object} responses.GetAllClasses
// @router 			/api/v1/classes [get]
func (h *ClassHandler) getAllClasses(c *fiber.Ctx) error {
	res, err := h.classRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllClasses{
		Items: *res,
	})
}
