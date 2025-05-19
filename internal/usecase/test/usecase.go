package test

import (
	"context"
	"diprec_api/internal/domain"
	"diprec_api/internal/repository/test"

	"go.uber.org/zap"
)

type testUsecase struct {
	repo   test.ITestRepository
	logger *zap.Logger
}

type ITestUsecase interface {
	Create(ctx context.Context, test *domain.Test, courseID uint) (*domain.Test, error)
	Get(ctx context.Context, courseID uint) ([]*domain.Test, error)
	GetByID(ctx context.Context, id uint) (*domain.Test, error)
	Update(ctx context.Context, test *domain.Test) (*domain.Test, error)
	Delete(ctx context.Context, id uint) error
	AttachQuestion(ctx context.Context, testID uint, questionID uint) error
	DetachQuestion(ctx context.Context, testID uint, questionID uint) error
}

func NewTestUsecase(repo test.ITestRepository, logger *zap.Logger) ITestUsecase {
	return &testUsecase{repo, logger.Named("TestUsecase")}
}

func (u *testUsecase) Create(ctx context.Context, test *domain.Test, courseID uint) (*domain.Test, error) {
	if err := u.repo.Create(ctx, test, courseID); err != nil {
		return nil, err
	}

	return test, nil
}

func (u *testUsecase) Get(ctx context.Context, courseID uint) ([]*domain.Test, error) {
	tests, err := u.repo.Get(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

func (u *testUsecase) GetByID(ctx context.Context, id uint) (*domain.Test, error) {
	test, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return test, nil
}

func (u *testUsecase) Update(ctx context.Context, test *domain.Test) (*domain.Test, error) {
	if err := u.repo.Update(ctx, test); err != nil {
		return nil, err
	}

	return test, nil
}

func (u *testUsecase) Delete(ctx context.Context, id uint) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (u *testUsecase) AttachQuestion(ctx context.Context, testID uint, questionID uint) error {
	if err := u.repo.AttachQuestion(ctx, testID, questionID); err != nil {
		return err
	}

	return nil
}

func (u *testUsecase) DetachQuestion(ctx context.Context, testID uint, questionID uint) error {
	if err := u.repo.DetachQuestion(ctx, testID, questionID); err != nil {
		return err
	}

	return nil
}
