package domain

import (
	"diprec_api/internal/pkg/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID       uint           `gorm:"primaryKey;autoIncrement"`
	Title    string         `gorm:"not null"`
	Type     string         `gorm:"not null"`
	Variants datatypes.JSON `gorm:"type:jsonb"`
	Answer   datatypes.JSON `gorm:"type:jsonb"`
	Tests    []Test         `gorm:"many2many:test_questions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type QuestionResponse struct {
	ID       uint                   `json:"id"`
	Title    string                 `json:"title"`
	Type     string                 `json:"type"`
	Variants map[string]interface{} `json:"variants"`
	Answer   map[string]interface{} `json:"answer"`
}

func (c *Question) ToQuestionResponse() QuestionResponse {
	return QuestionResponse{
		ID:       c.ID,
		Title:    c.Title,
		Type:     c.Type,
		Variants: utils.ParseJSONToMap(c.Variants),
		Answer:   utils.ParseJSONToMap(c.Answer),
	}
}

func ToQuestionsResponse(questions []*Question) []QuestionResponse {
	responses := make([]QuestionResponse, len(questions))
	for i, question := range questions {
		responses[i] = QuestionResponse{
			ID:       question.ID,
			Title:    question.Title,
			Type:     question.Type,
			Variants: utils.ParseJSONToMap(question.Variants),
			Answer:   utils.ParseJSONToMap(question.Answer),
		}
	}

	return responses
}
