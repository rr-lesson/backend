package requests

type CreateSubject struct {
	ClassId uint   `json:"class_id"`
	Name    string `json:"name"`
} //@name CreateSubjectReq
