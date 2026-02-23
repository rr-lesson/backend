package dto

import "backend/internal/domains"

type UserDTO struct {
	Data domains.User `json:"data"`
} // @name UserDTO
