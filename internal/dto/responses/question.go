package responses

import (
	"backend/internal/domains"
	"backend/internal/dto"
)

type CreateQuestion struct {
	Question domains.Question `json:"question"`
} // @name CreateQuestionRes

type GetAllQuestions struct {
	Items []dto.QuestionDTO `json:"items"`
} // @name GetAllQuestionsRes

type GetQuestion struct {
	Question dto.QuestionDTO `json:"question"`
} // @name GetQuestionRes
