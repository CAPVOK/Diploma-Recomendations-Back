package question

import (
	"context"
	"diprec_api/internal/domain"
	"diprec_api/internal/pkg/utils"
	"diprec_api/internal/repository/question"
	"go.uber.org/zap"
)

type questionUsecase struct {
	repo   question.IQuestionRepository
	logger *zap.Logger
}

type IQuestionUsecase interface {
	Create(ctx context.Context, question *domain.Question) (*domain.Question, error)
	GetByID(ctx context.Context, id uint) (*domain.Question, error)
	Update(ctx context.Context, question *domain.Question) (*domain.Question, error)
	Delete(ctx context.Context, id uint) error
	Check(ctx context.Context, id uint, answer interface{}) (*domain.QuestionAnswer, error)
}

func NewQuestionUsecase(repo question.IQuestionRepository, logger *zap.Logger) IQuestionUsecase {
	return &questionUsecase{repo, logger}
}

func (u *questionUsecase) Create(ctx context.Context, question *domain.Question) (*domain.Question, error) {
	if err := u.repo.Create(ctx, question); err != nil {
		return nil, err
	}

	return question, nil
}

func (u *questionUsecase) GetByID(ctx context.Context, id uint) (*domain.Question, error) {
	question, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (u *questionUsecase) Update(ctx context.Context, question *domain.Question) (*domain.Question, error) {
	if err := u.repo.Update(ctx, question); err != nil {
		return nil, err
	}

	return question, nil
}

func (u *questionUsecase) Delete(ctx context.Context, id uint) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (u *questionUsecase) Check(ctx context.Context, id uint, answer interface{}) (*domain.QuestionAnswer, error) {
	question, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	isCorrect := question.CheckAnswer(answer)
	return &domain.QuestionAnswer{
		IsCorrect: isCorrect,
		Message:   utils.GenerateFeedbackMessage(isCorrect),
		Answer:    question.Answer,
	}, nil
}
