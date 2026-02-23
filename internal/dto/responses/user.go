package responses

import "backend/internal/dto"

type GetAllUsers struct {
	Items []dto.UserDTO `json:"items"`
} // @name GetAllUsersRes

type GetCurrentUser struct {
	User dto.UserDTO `json:"user"`
} // @name GetCurrentUserRes

type UpdateCurrentUser struct {
	User dto.UserDTO `json:"user"`
} // @name UpdateCurrentUserRes
