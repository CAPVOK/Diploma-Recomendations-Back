package domain

import (
	"time"

	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null;unique"`
	Description string
	Deadline    time.Time   `gorm:"not null"`
	Courses     []*Course   `gorm:"many2many:course_tests;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Questions   []*Question `gorm:"many2many:test_questions;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}

type TestResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TestResponseWithQuestions struct {
	TestResponse
	Questions []QuestionResponse `json:"questions"`
}

func (c *Test) ToTestResponse() TestResponse {
	return TestResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Deadline:    c.Deadline,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (c *Test) ToTestResponseWithQuestions() TestResponseWithQuestions {
	return TestResponseWithQuestions{
		TestResponse: c.ToTestResponse(),
		Questions:    ToQuestionsResponse(c.Questions),
	}
}

func ToTestsResponse(test []*Test) []TestResponse {
	response := make([]TestResponse, len(test))
	for i, test := range test {
		response[i] = test.ToTestResponse()
	}

	return response
}
