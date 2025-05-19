package question

import (
	"context"
	"diprec_api/internal/domain"
	"diprec_api/internal/infrastructure/kafka"
	"diprec_api/internal/pkg/utils"
	"diprec_api/internal/repository/question"

	"go.uber.org/zap"
	"strconv"
	"time"
)

type questionUsecase struct {
	repo     question.IQuestionRepository
	producer kafka.IKafkaProducer
	logger   *zap.Logger
}

type IQuestionUsecase interface {
	Create(ctx context.Context, question *domain.Question) (*domain.Question, error)
	GetAll(ctx context.Context) ([]*domain.Question, error)
	GetByID(ctx context.Context, id uint) (*domain.Question, error)
	Update(ctx context.Context, question *domain.Question) (*domain.Question, error)
	Delete(ctx context.Context, id uint) error
	Check(ctx context.Context, id, userID uint, answer interface{}) (*domain.QuestionAnswer, error)
}

func NewQuestionUsecase(repo question.IQuestionRepository, producer kafka.IKafkaProducer, logger *zap.Logger) IQuestionUsecase {
	return &questionUsecase{repo, producer, logger}
}

func (u *questionUsecase) Create(ctx context.Context, question *domain.Question) (*domain.Question, error) {
	if err := u.repo.Create(ctx, question); err != nil {
		return nil, err
	}

	return question, nil
}

func (u *questionUsecase) GetAll(ctx context.Context) ([]*domain.Question, error) {
	questions, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return questions, nil
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

func (u *questionUsecase) Check(ctx context.Context, id, userID uint, answer interface{}) (*domain.QuestionAnswer, error) {
	question, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	isCorrect := question.CheckAnswer(answer)

	userAnswer := &domain.UserAnswer{
		QuestionID: question.ID,
		Title:      question.Title,
		Type:       question.Type.String(),
		Variants:   question.Variants,
		Answer:     question.Answer,
		IsCorrect:  isCorrect,
		UserID:     userID,
		Timestamp:  time.Now(),
	}

	_ = u.producer.Send(ctx, strconv.Itoa(int(userID)), userAnswer)

	return &domain.QuestionAnswer{
		IsCorrect: isCorrect,
		Message:   utils.GenerateFeedbackMessage(isCorrect),
		Answer:    question.Answer,
	}, nil
}
