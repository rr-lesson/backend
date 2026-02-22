package handlers

import (
	"backend/internal/domains"
	"backend/internal/dto/requests"
	"backend/internal/dto/responses"
	"backend/internal/repositories"
	"backend/pkg/auth"
	"backend/pkg/utils"
	"encoding/json"

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
	g0 := router.Group("/questions").Use(auth.AuthMiddleware())
	g0.Post("/", h.createQuestion)
	g0.Get("/", h.getAllQuestions)
	g0.Get("/:questionId", h.getQuestion)
}

// @id 					CreateQuestion
// @tags 				question
// @accept 			multipart/form-data
// @produce 		json
// @param 			images formData []file false "images"
// @param 			body formData string true "body"
// @success 		200 {object} responses.CreateQuestion
// @router 			/api/v1/questions [post]
func (h *QuestionHandler) createQuestion(c *fiber.Ctx) error {
	session := h.authHelper.GetCurrentSession(c)
	if session == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Error{
			Message: "Anda tidak memiliki akses untuk melakukan aksi ini!",
		})
	}

	formData, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	body := formData.Value["body"][0]
	images := formData.File["images"]

	var req requests.CreateQuestion
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.questionRepo.Create(
		c.Context(),
		domains.Question{
			UserId:    session.UserId,
			SubjectId: req.SubjectId,
			Question:  req.Question,
		},
		images,
	)
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
// @param 			keyword query string false "keyword"
// @param 			includes query []string false "includes" Enums(user, subject, class)
// @success 		200 {object} responses.GetAllQuestions
// @router 			/api/v1/questions [get]
func (h *QuestionHandler) getAllQuestions(c *fiber.Ctx) error {
	includes := utils.ParseIncludes(c)

	res, err := h.questionRepo.GetAll(repositories.QuestionFilter{
		Keyword:        c.Query("keyword"),
		IncludeUser:    includes["user"],
		IncludeSubject: includes["subject"],
		IncludeClass:   includes["class"],
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllQuestions{
		Items: *res,
	})
}

// @id 					GetQuestion
// @tags 				question
// @accept 			json
// @produce 		json
// @param 			questionId path int true "questionId"
// @param 			includes query []string false "includes" Enums(user, subject, class, attachments)
// @success 		200 {object} responses.GetQuestion
// @router 			/api/v1/questions/{questionId} [get]
func (h *QuestionHandler) getQuestion(c *fiber.Ctx) error {
	questionId, err := c.ParamsInt("questionId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	includes := utils.ParseIncludes(c)

	res, err := h.questionRepo.Get(repositories.QuestionFilter{
		QuestionId:         uint(questionId),
		IncludeUser:        includes["user"],
		IncludeSubject:     includes["subject"],
		IncludeClass:       includes["class"],
		IncludeAttachments: includes["attachments"],
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetQuestion{
		Question: *res,
	})
}
