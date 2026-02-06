package requests

type CreateLesson struct {
	SubjectId uint   `json:"subject_id"`
	Title     string `json:"title"`
} //@name CreateLessonReq
