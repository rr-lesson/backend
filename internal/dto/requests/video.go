package requests

type CreateVideo struct {
	LessonId    uint   `json:"lesson_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FilePath    string `json:"file_path"`
} //@name CreateVideoReq
