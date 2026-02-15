package handlers

import (
	"backend/internal/domains"
	"backend/internal/dto/requests"
	"backend/internal/dto/responses"
	"backend/internal/repositories"
	"backend/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

type QuestionHandler struct {
	authRepo     *repositories.AuthRepository
	questionRepo *repositories.QuestionRepository
	authHelper   *auth.AuthHelper
}

func NewQuestionHandler(
	authRepo *repositories.AuthRepository,
	questionRepo *repositories.QuestionRepository,
	authHelper *auth.AuthHelper,
) *QuestionHandler {
	return &QuestionHandler{
		authRepo:     authRepo,
		questionRepo: questionRepo,
		authHelper:   authHelper,
	}
}

func (h *QuestionHandler) RegisterRoutes(router fiber.Router) {
	g0 := router.Group("/questions")
	g0.Post("/", h.createQuestion)
	g0.Get("/", h.getAllQuestions)
}

// @id 					CreateQuestion
// @tags 				question
// @accept 			json
// @produce 		json
// @param 			body body requests.CreateQuestion true "body"
// @success 		200 {object} responses.CreateQuestion
// @router 			/api/v1/questions [post]
func (h *QuestionHandler) createQuestion(c *fiber.Ctx) error {
	session := h.authHelper.GetCurrentSession(c)
	if session == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Error{
			Message: "Anda tidak memiliki akses untuk melakukan aksi ini!",
		})
	}

	var req requests.CreateQuestion
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.questionRepo.Create(domains.Question{
		UserId:    session.UserId,
		SubjectId: req.SubjectId,
		Question:  req.Question,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.CreateQuestion{
		Question: *res,
	})
}

// @id 					GetAllQuestions
// @tags 				question
// @accept 			json
// @produce 		json
// @success 		200 {object} responses.GetAllQuestions
// @router 			/api/v1/questions [get]
func (h *QuestionHandler) getAllQuestions(c *fiber.Ctx) error {
	res, err := h.questionRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllQuestions{
		Questions: *res,
	})
}
