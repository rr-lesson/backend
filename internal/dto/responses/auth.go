package responses

import "backend/internal/domains"

type Login struct {
	User domains.User `json:"user"`
} // @name LoginRes

type Register struct {
	User domains.User `json:"user"`
} // @name RegisterRes

type Logout struct {
	Message string `json:"message"`
} // @name LogoutRes
