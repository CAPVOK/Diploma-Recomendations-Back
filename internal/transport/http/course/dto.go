package course

import (
	"duolingo_api/internal/domain"
	"time"
)

type CreateCourseDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCourseDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CourseResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ToCourseResponse(course *domain.Course) CourseResponse {
	return CourseResponse{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}
}
