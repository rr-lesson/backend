package requests

type CreateVideo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FilePath    string `json:"file_path"`
} //@name CreateVideoReq
