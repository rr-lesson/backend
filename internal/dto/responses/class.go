package responses

import (
	"backend/internal/domains"
	"backend/internal/dto"
)

type CreateClass struct {
	Class domains.Class `json:"class"`
} //@name CreateClassRes

type GetAllClasses struct {
	Items []dto.ClassDTO `json:"items"`
} //@name GetAllClassesRes
