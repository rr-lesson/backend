package responses

import "backend/internal/domains"

type CreateClass struct {
	Class domains.Class `json:"class"`
} //@name CreateClassRes

type GetAllClasses struct {
	Classes []domains.Class `json:"classes"`
} //@name GetAllClassesRes
