package handlers

import (
	"backend/internal/domains"
	"backend/internal/dto/requests"
	"backend/internal/dto/responses"
	"backend/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

type LessonHandler struct {
	lessonRepo *repositories.LessonRepository
}

func NewLessonHandler(
	lessonRepo *repositories.LessonRepository,
) *LessonHandler {
	return &LessonHandler{
		lessonRepo: lessonRepo,
	}
}

func (h *LessonHandler) RegisterRoutes(router fiber.Router) {
	g0 := router.Group("/classes/:classId/subjects/:subjectId/lessons")
	g0.Get("/", h.getAllLessonsBySubjectId)

	g1 := router.Group("/classes/subjects/lessons")
	g1.Get("/", h.getAllLessonWithClassSubject)

	g2 := router.Group("/lessons")
	g2.Post("/", h.createLesson)
	g2.Get("/", h.getAllLessons)
}

// @id 					CreateLesson
// @tags 				lesson
// @accept 			json
// @produce 		json
// @param 			body body requests.CreateLesson true "body"
// @success 		200 {object} responses.CreateLesson
// @router 			/api/v1/lessons [post]
func (h *LessonHandler) createLesson(c *fiber.Ctx) error {
	var req requests.CreateLesson
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.lessonRepo.Create(domains.Lesson{
		SubjectId: req.SubjectId,
		Title:     req.Title,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.CreateLesson{
		Lesson: *res,
	})
}

// @id 					GetAllLessons
// @tags 				lesson
// @accept 			json
// @produce 		json
// @param 			classId query int false "classId"
// @param 			subjectId query int false "subjectId"
// @success 		200 {object} responses.GetAllLessons
// @router 			/api/v1/lessons [get]
func (h *LessonHandler) getAllLessons(c *fiber.Ctx) error {
	classId := c.QueryInt("classId", 0)
	subjectId := c.QueryInt("subjectId", 0)

	res, err := h.lessonRepo.GetAll(repositories.LessonFilter{
		ClassId:   uint(classId),
		SubjectId: uint(subjectId),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllLessons{
		Lessons: *res,
	})
}

// @id 					GetAllLessonWithClassSubject
// @tags 				lesson
// @accept 			json
// @produce 		json
// @success 		200 {object} responses.GetAllLessonWithClassSubject
// @router 			/api/v1/classes/subjects/lessons [get]
func (h *LessonHandler) getAllLessonWithClassSubject(c *fiber.Ctx) error {
	res, err := h.lessonRepo.GetAllWithClassSubject()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllLessonWithClassSubject{
		Lessons: *res,
	})
}

// @deprecated
// @id 					GetAllLessonsBySubjectId
// @tags 				lesson
// @accept 			json
// @produce 		json
// @param 			classId path int true "classId"
// @param 			subjectId path int true "subjectId"
// @success 		200 {object} responses.GetAllLessonsBySubjectId
// @router 			/api/v1/classes/{classId}/subjects/{subjectId}/lessons [get]
func (h *LessonHandler) getAllLessonsBySubjectId(c *fiber.Ctx) error {
	subjectId, err := c.ParamsInt("subjectId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.lessonRepo.GetAllBySubjectId(uint(subjectId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllLessonsBySubjectId{
		Lessons: *res,
	})
}
