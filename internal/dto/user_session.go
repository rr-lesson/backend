package dto

import "backend/internal/domains"

type UserSessionDTO struct {
	Data domains.UserSession `json:"data"`
	User domains.User        `json:"user"`
} // @name UserSessionDTO
