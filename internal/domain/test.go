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
	Assignee    Assignee    `json:"assignee" gorm:"type:varchar(20);not null;assignee IN ('TEACHER', 'RECOMMENDATION');default:'TEACHER'"`
	Deadline    time.Time   `gorm:"not null"`
	Courses     []*Course   `gorm:"many2many:course_tests;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Questions   []*Question `gorm:"many2many:test_questions;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}

type TestResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Assignee    string    `json:"assignee"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TestResponseWithQuestions struct {
	TestResponse
	Questions []QuestionResponse `json:"questions"`
}

type Assignee string

const (
	Teacher        Assignee = "TEACHER"
	Recommendation Assignee = "RECOMMENDATION"
)

func (a Assignee) String() string {
	return string(a)
}

func (c *Test) ToTestResponse() TestResponse {
	return TestResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Assignee:    c.Assignee.String(),
		Deadline:    c.Deadline,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (c *Test) ToTestResponseWithQuestions(isTeacher bool) TestResponseWithQuestions {
	return TestResponseWithQuestions{
		TestResponse: c.ToTestResponse(),
		Questions:    ToQuestionsResponse(c.Questions, isTeacher),
	}
}

func ToTestsResponse(test []*Test) []TestResponse {
	response := make([]TestResponse, len(test))
	for i, test := range test {
		response[i] = test.ToTestResponse()
	}

	return response
}
