package question

type CreateQuestionDTO struct {
	Title    string                 `json:"title"`
	Type     string                 `json:"type"`
	Variants map[string]interface{} `json:"variants"`
	Answer   map[string]interface{} `json:"answer"`
}

type UpdateQuestionDTO struct {
	Title    string                 `json:"title"`
	Type     string                 `json:"type"`
	Variants map[string]interface{} `json:"variants"`
	Answer   map[string]interface{} `json:"answer"`
}
