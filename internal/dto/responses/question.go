package responses

import "backend/internal/domains"

type CreateQuestion struct {
	Question domains.Question `json:"question"`
} // @name CreateQuestionRes

type GetAllQuestions struct {
	Questions []domains.Question `json:"questions"`
} // @name GetAllQuestionsRes
