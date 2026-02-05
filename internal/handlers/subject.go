package handlers

import (
	"backend/internal/domains"
	"backend/internal/dto/requests"
	"backend/internal/dto/responses"
	"backend/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

type SubjectHandler struct {
	subjectRepo *repositories.SubjectRepository
}

func NewSubjectHandler(
	subjectRepo *repositories.SubjectRepository,
) *SubjectHandler {
	return &SubjectHandler{
		subjectRepo: subjectRepo,
	}
}

func (h *SubjectHandler) RegisterRoutes(router fiber.Router) {
	g0 := router.Group("/subjects")
	g0.Post("/", h.createSubject)
	g0.Get("/", h.getAllSubjects)
	g0.Get("/details", h.getAllSubjectDetails)
}

// @id 					CreateSubject
// @tags 				subject
// @accept 			json
// @produce 		json
// @param 			body body requests.CreateSubject true "body"
// @success 		200 {object} responses.CreateSubject
// @router 			/api/v1/subjects [post]
func (h *SubjectHandler) createSubject(c *fiber.Ctx) error {
	var req requests.CreateSubject
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.subjectRepo.Create(domains.Subject{
		ClassId: req.ClassId,
		Name:    req.Name,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.CreateSubject{
		Subject: *res,
	})
}

// @id 					GetAllSubjects
// @tags 				subject
// @accept 			json
// @produce 		json
// @param 			classId query int false "classId"
// @success 		200 {object} responses.GetAllSubjects
// @router 			/api/v1/subjects [get]
func (h *SubjectHandler) getAllSubjects(c *fiber.Ctx) error {
	classId := c.QueryInt("classId", 0)

	res, err := h.subjectRepo.GetAll(uint(classId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllSubjects{
		Subjects: *res,
	})
}

// @id 					GetAllSubjectDetails
// @tags 				subject
// @accept 			json
// @produce 		json
// @success 		200 {object} responses.GetAllSubjectDetails
// @router 			/api/v1/subjects/details [get]
func (h *SubjectHandler) getAllSubjectDetails(c *fiber.Ctx) error {
	res, err := h.subjectRepo.GetAllDetails()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllSubjectDetails{
		Subjects: *res,
	})
}
