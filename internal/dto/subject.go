package dto

import "backend/internal/domains"

type SubjectDetail struct {
	Subject domains.Subject `json:"subject"`
	Class   domains.Class   `json:"class"`
} // @name SubjectDetail
