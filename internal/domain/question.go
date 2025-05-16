package domain

import (
	"diprec_api/internal/pkg/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID       uint           `gorm:"primaryKey;autoIncrement"`
	Title    string         `gorm:"not null,unique"`
	Type     Type           `gorm:"type:varchar(20);not null;type IN ('SINGLE', 'MULTIPLE', 'TEXT', 'NUMBER');default:'SINGLE'"`
	Variants datatypes.JSON `gorm:"type:jsonb"`
	Answer   datatypes.JSON `gorm:"type:jsonb"`
	Tests    []Test         `gorm:"many2many:test_questions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Type string

const (
	Single   Type = "SINGLE"
	Multiple Type = "MULTIPLE"
	Text     Type = "TEXT"
	Number   Type = "NUMBER"
)

func (t Type) String() string {
	return string(t)
}

type QuestionResponse struct {
	ID       uint                   `json:"id"`
	Title    string                 `json:"title"`
	Type     string                 `json:"type" enums:"SINGLE,MULTIPLE,TEXT,NUMBER" example:"SINGLE"`
	Variants map[string]interface{} `json:"variants"`
	Answer   map[string]interface{} `json:"answer"`
}

func (c *Question) ToQuestionResponse() QuestionResponse {
	return QuestionResponse{
		ID:       c.ID,
		Title:    c.Title,
		Type:     c.Type.String(),
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
			Type:     question.Type.String(),
			Variants: utils.ParseJSONToMap(question.Variants),
			Answer:   utils.ParseJSONToMap(question.Answer),
		}
	}

	return responses
}
