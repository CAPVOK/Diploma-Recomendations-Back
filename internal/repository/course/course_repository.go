package course

import (
	"context"
	"duolingo_api/internal/domain"
	"duolingo_api/internal/pkg/validator"
	"errors"
	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

type ICourseRepository interface {
	Create(ctx context.Context, course *domain.Course) error
	Get(ctx context.Context) ([]*domain.Course, error)
	GetByID(ctx context.Context, id uint) (*domain.Course, error)
	Update(ctx context.Context, course *domain.Course) error
	Delete(ctx context.Context, id uint) error
}

func NewCourseRepository(db *gorm.DB) ICourseRepository { return &courseRepository{db: db} }

func (r *courseRepository) Create(ctx context.Context, course *domain.Course) error {
	err := r.db.Create(course).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *courseRepository) Get(ctx context.Context) ([]*domain.Course, error) {
	var courses []*domain.Course

	err := r.db.Find(&courses).Error
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *courseRepository) GetByID(ctx context.Context, id uint) (*domain.Course, error) {
	var course *domain.Course

	err := r.db.First(&course, id).Error
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (r *courseRepository) Update(ctx context.Context, course *domain.Course) error {
	updates := validator.BuildUpdates(course)

	result := r.db.Model(&domain.Course{}).Where("id =?", course.ID).Updates(updates).First(&course)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.ErrCourseNotFound
		}

		return result.Error
	}

	return nil
}

func (r *courseRepository) Delete(ctx context.Context, id uint) error {
	err := r.db.Delete(&domain.Course{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
