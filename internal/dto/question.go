package dto

import "backend/internal/domains"

type QuestionDTO struct {
	User        domains.User                 `json:"user"`
	Subject     domains.Subject              `json:"subject"`
	Class       domains.Class                `json:"class"`
	Data        domains.Question             `json:"data"`
	Attachments []domains.QuestionAttachment `json:"attachments"`
} // @name QuestionDTO
