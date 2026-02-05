package responses

import (
	"backend/internal/domains"
	"backend/internal/dto"
)

type CreateLesson struct {
	Lesson domains.Lesson `json:"lesson"`
} //@name CreateLessonRes

type GetAllLessons struct {
	Lessons []domains.Lesson `json:"lessons"`
} //@name GetAllLessonsRes

type GetAllLessonWithClassSubject struct {
	Lessons []dto.LessonClassSubject `json:"lessons"`
} //@name GetAllLessonWithClassSubjectRes

type GetAllLessonsBySubjectId struct {
	Lessons []domains.Lesson `json:"lessons"`
} //@name GetAllLessonsBySubjectIdRes
