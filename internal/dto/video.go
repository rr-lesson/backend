package dto

import "backend/internal/domains"

type VideoDetail struct {
	Video   domains.Video   `json:"video"`
	Lesson  domains.Lesson  `json:"lesson"`
	Subject domains.Subject `json:"subject"`
	Class   domains.Class   `json:"class"`
} //@name VideoDetail
