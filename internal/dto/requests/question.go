package requests

type CreateQuestion struct {
	SubjectId uint   `json:"subject_id"`
	Question  string `json:"question"`
} // @name CreateQuestionReq
