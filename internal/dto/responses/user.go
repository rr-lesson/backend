package responses

import "backend/internal/dto"

type GetAllUsers struct {
	Items []dto.UserDTO `json:"items"`
} // @name GetAllUsersRes
