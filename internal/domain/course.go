package domain

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null"`
	Description string
	Users       []User `gorm:"many2many:user_courses;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tests       []Test `gorm:"many2many:course_tests;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}

type CourseResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (c *Course) ToCourseResponse() CourseResponse {
	return CourseResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func ToCoursesResponse(courses []*Course) []CourseResponse {
	responses := make([]CourseResponse, len(courses))
	for i, course := range courses {
		responses[i] = course.ToCourseResponse()
	}
	return responses
}
