package dto

import "backend/internal/domains"

type LessonClassSubject struct {
	Lesson  domains.Lesson  `json:"lesson"`
	Class   domains.Class   `json:"class"`
	Subject domains.Subject `json:"subject"`
} //@name LessonClassSubject
