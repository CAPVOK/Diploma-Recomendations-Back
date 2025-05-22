package test

import "time"

type CreateTestDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type UpdateTestDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type AttachQuestionDTO struct {
	QuestionID int `json:"questionId"`
}

type FinishTestDTO struct {
	Progress uint `json:"progress"`
}
