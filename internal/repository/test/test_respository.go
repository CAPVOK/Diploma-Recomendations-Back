package test

import (
	"context"
	"diprec_api/internal/domain"
	"diprec_api/internal/pkg/validator"
	"errors"
	"gorm.io/gorm"
)

type testRepository struct {
	db *gorm.DB
}

type ITestRepository interface {
	Create(ctx context.Context, test *domain.Test, courseID uint) error
	Get(ctx context.Context, courseID uint) ([]*domain.Test, error)
	GetByID(ctx context.Context, id uint) (*domain.Test, error)
	Update(ctx context.Context, test *domain.Test) error
	Delete(ctx context.Context, id uint) error
}

func NewTestRepository(db *gorm.DB) ITestRepository { return &testRepository{db: db} }

func (r *testRepository) Create(ctx context.Context, test *domain.Test, courseID uint) error {
	if err := r.db.Create(test).Error; err != nil {
		return err
	}

	return r.db.Model(test).Association("Courses").Append(&domain.Course{ID: courseID})
}

func (r *testRepository) Get(ctx context.Context, courseID uint) ([]*domain.Test, error) {
	var course domain.Course

	err := r.db.Preload("Tests").First(&course, courseID).Error
	if err != nil {
		return nil, err
	}

	return course.Tests, nil
}

func (r *testRepository) GetByID(ctx context.Context, id uint) (*domain.Test, error) {
	var test domain.Test

	err := r.db.Where("id = ?", id).First(&test).Error
	if err != nil {
		return nil, err
	}

	return &test, nil
}

func (r *testRepository) Update(ctx context.Context, test *domain.Test) error {
	updates := validator.BuildUpdates(test)

	result := r.db.Model(&domain.Test{}).Where("id = ?", test.ID).Updates(updates).First(&test)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.ErrTestNotFound
		}

		return result.Error
	}

	return nil
}

func (r *testRepository) Delete(ctx context.Context, id uint) error {
	var test domain.Test

	if err := r.db.First(&test, id).Error; err != nil {
		return err
	}

	if err := r.db.Model(&test).Association("Courses").Clear(); err != nil {
		return err
	}

	return r.db.Delete(&test).Error
}
