package responses

import (
	"backend/internal/domains"
	"backend/internal/dto"
)

type CreateSubject struct {
	Subject domains.Subject `json:"subject"`
} // @name CreateSubjectRes

type GetAllSubjects struct {
	Subjects []domains.Subject `json:"subjects"`
} // @name GetAllSubjectsRes

type GetAllSubjectDetails struct {
	Subjects []dto.SubjectDetail `json:"subjects"`
} // @name GetAllSubjectDetailsRes
